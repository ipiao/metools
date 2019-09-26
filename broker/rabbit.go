package broker

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipiao/meim/log"
	"time"
)

type RPCHandler func([]byte) []byte

var (
	timeoutByte  = []byte("timeout")
	ErrorTimeout = errors.New("timeout")
)

const (
	ChannelPub = 1 << iota
	ChannelSub
	ChannelRPC
	ChannelRPCServer
	AllChannel = ChannelPub | ChannelSub | ChannelRPC | ChannelRPCServer
)

// rabbit配置
type RabbitMQConfig struct {
	Node         int
	Url          string
	ExchangeName string
	ExchangeKind string
	ChanSize     int
	Channels     uint64
	RPCTimeout   time.Duration
	SendTimeout  time.Duration
	QueuePrefix  string // 队列前缀
}

func (cfg *RabbitMQConfig) init() {
	if cfg.ExchangeName == "" {
		panic("exchange name not set")
	}
	if cfg.ExchangeKind != "direct" && cfg.ExchangeKind != "fanout" && cfg.ExchangeKind != "topic" {
		cfg.ExchangeKind = "direct"
	}
	if cfg.RPCTimeout == 0 {
		cfg.RPCTimeout = time.Second * 5
	}
	if cfg.ChanSize == 0 {
		cfg.ChanSize = 512
	}
	if cfg.QueuePrefix == "" {
		cfg.QueuePrefix = "message"
	}
	if cfg.Node <= 0 {
		panic("cfg node must be greater than 0")
	}
}

// 用户rpc操作的请求结构
type request struct {
	node int
	body []byte
	ret  chan []byte
}

// 一个rabbit完整的Broker
type RabbitMQ struct {
	cancel         context.CancelFunc
	cfg            *RabbitMQConfig
	pubMessageChan chan *request // pub message
	rpcRequestChan chan *request // rpc message
	subMessageChan chan []byte   // sub message
	rpcHandler     RPCHandler
}

// 新建rabbot通道,参数需要给定
func NewRabbitMQ(cfg *RabbitMQConfig, rpcHandler RPCHandler) *RabbitMQ {
	cfg.init()

	ctx, done := context.WithCancel(context.Background())
	rb := &RabbitMQ{
		cancel:     done,
		cfg:        cfg,
		rpcHandler: rpcHandler,
	}

	if cfg.Channels&ChannelSub != 0 {
		go func() {
			rb.subMessageChan = make(chan []byte, cfg.ChanSize)
			rb.subscribe(redial(ctx, cfg.Url, rb.cfg.ExchangeName, rb.cfg.ExchangeKind))
			done()
		}()
	}

	if cfg.Channels&ChannelPub != 0 {
		go func() {
			rb.pubMessageChan = make(chan *request, cfg.ChanSize)
			rb.publish(redial(ctx, cfg.Url, rb.cfg.ExchangeName, rb.cfg.ExchangeKind))
			done()
		}()
	}

	if cfg.Channels&ChannelRPCServer != 0 {
		if rb.rpcHandler == nil {
			panic(" rpcHandler is not set")
		}
		go func() {
			rb.rpcServer(redial(ctx, cfg.Url, rb.cfg.ExchangeName, rb.cfg.ExchangeKind))
			done()
		}()
	}
	if cfg.Channels&ChannelRPC != 0 {
		go func() {
			rb.rpcRequestChan = make(chan *request, cfg.ChanSize)
			rb.rpc(redial(ctx, cfg.Url, rb.cfg.ExchangeName, rb.cfg.ExchangeKind))
			done()
		}()
	}
	return rb
}

// 发送消息
func (rb *RabbitMQ) publish(sessions chan chan session) {
	for session := range sessions {
		pub := <-session
		if !pub.connected() {
			log.Warnf("[rabbit] session not connected")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		var (
			reading = rb.pubMessageChan
			err     error
		)

		log.Debug("[rabbit] publishing...")
	Publish:
		for {
			select {
			case req, ok := <-reading:
				if !ok {
					return
				}
				routineKey := rb.getRoutingKey(req.node)
				err = pub.Publish(rb.cfg.ExchangeName, routineKey, false, false, amqp.Publishing{
					Body: req.body,
				})
				if err != nil {
					log.Errorf("[rabbit] can not publish message: %v", err)
					reading <- req
					pub.close()
					break Publish
				}
			}
		}
	}
}

// 订阅消息
func (rb *RabbitMQ) subscribe(sessions chan chan session) {

	queue := rb.getQueueName()
	for session := range sessions {
		sub := <-session
		if !sub.connected() {
			log.Warnf("[rabbit] session not connected")
			time.Sleep(time.Millisecond * 10)
			continue
		}

		// 去除排他性
		if _, err := sub.QueueDeclare(queue, true, false, false, false, nil); err != nil {
			log.Errorf("[rabbit] cannot consume from exclusive queue: %q, %v", queue, err)
			sub.close()
			continue
		}

		if err := sub.QueueBind(queue, rb.getBindKey(), rb.cfg.ExchangeName, false, nil); err != nil {
			log.Errorf("[rabbit] cannot consume without a binding to exchange: %q, %v", rb.cfg.ExchangeName, err)
			sub.close()
			continue
		}

		deliveries, err := sub.Consume(queue, "", true, true, false, false, nil)
		if err != nil {
			log.Errorf("[rabbit]  cannot consume from: %q, %v", queue, err)
			sub.close()
			continue
		}

		log.Debug("[rabbit] subscribed...")
		for msg := range deliveries {
			rb.subMessageChan <- msg.Body
			// sub.Ack(msg.DeliveryTag, false)
		}
	}
}

// 开启rpc服务
func (rb *RabbitMQ) rpcServer(sessions chan chan session) {
	rpcQueueName := rb.getRpcQueueName()
	for session := range sessions {
		rpc := <-session
		if !rpc.connected() {
			log.Warnf("[rabbit] session not connected")
			time.Sleep(time.Millisecond * 10)
			continue
		}

		if _, err := rpc.QueueDeclare(rpcQueueName, true, false, false, false, nil); err != nil {
			log.Errorf("[rabbit] cannot consume from exclusive queue: %q, %v", rpcQueueName, err)
			rpc.close()
			continue
		}

		if err := rpc.QueueBind(rpcQueueName, rb.getRpcBindKey(), rb.cfg.ExchangeName, false, nil); err != nil {
			log.Errorf("[rabbit] cannot consume without a binding to exchange: %q, %v", rb.cfg.ExchangeName, err)
			rpc.close()
			continue
		}

		msgs, err := rpc.Consume(rpcQueueName, "", true, false, false, false, nil)
		if err != nil {
			log.Errorf("[rabbit] cannot consume from exclusive queue: %q, %v", rpcQueueName, err)
			rpc.close()
			continue
		}

		log.Debug("[rabbit] serving rpc...")
		for d := range msgs {
			err = rpc.Publish(rb.cfg.ExchangeName, d.ReplyTo, false, false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          rb.rpcHandler(d.Body),
				})
		}
	}
}

// 发送消息
func (rb *RabbitMQ) rpc(sessions chan chan session) {
	for session := range sessions {
		log.Debug("[rabbit] rpc...")

		rpc := <-session
		var reqs = rb.rpcRequestChan
	PUBLISH:
		for {
			select {
			case req := <-reqs:
				corrId := uuid.New().String()
				// 接收rpc服务返回队列
				q, err := rpc.QueueDeclare(corrId, false, true, true, false, nil)
				if err != nil {
					log.Errorf("[rabbit] cannot consume from exclusive queue: %q, %v", q.Name, err)
					rpc.close()
					reqs <- req
					break PUBLISH
				}

				if err := rpc.QueueBind(q.Name, q.Name, rb.cfg.ExchangeName, false, nil); err != nil {
					log.Errorf("[rabbit] cannot consume without a binding to exchange: %q, %v", rb.cfg.ExchangeName, err)
					rpc.close()
					reqs <- req
					break PUBLISH
				}

				msgs, err := rpc.Consume(q.Name, "", true, true, false, false, nil)
				if err != nil {
					log.Errorf("[rabbit] cannot consume from: %q, %v", rb.getRpcQueueName(), err)
					rpc.close()
					reqs <- req
					break PUBLISH
				}

				routineKey := rb.getRpcRoutingKey(req.node)
				err = rpc.Publish(rb.cfg.ExchangeName, routineKey, false, false,
					amqp.Publishing{
						ContentType:   "text/plain",
						CorrelationId: corrId,
						ReplyTo:       q.Name,
						Body:          req.body,
					})
				if err != nil {
					log.Errorf("[rabbit] can not publish message: %v", err)
					reqs <- req
					rpc.close()
					break PUBLISH
				}
				go func() {
					for {
						select {
						case d := <-msgs:
							if corrId == d.CorrelationId {
								req.ret <- d.Body
								return
							}
						case <-time.After(rb.cfg.RPCTimeout):
							req.ret <- timeoutByte
							return
						}
					}
				}()
			}
		}
	}
}

// 生成队列名
func (rb *RabbitMQ) getQueueName() string {
	return fmt.Sprintf("%s_%d", rb.cfg.QueuePrefix, rb.cfg.Node)
}

// 队列绑定key
func (rb *RabbitMQ) getBindKey() string {
	return fmt.Sprintf("%s.%d.*", rb.cfg.QueuePrefix, rb.cfg.Node)
}

// rpc队列名
func (rb *RabbitMQ) getRpcQueueName() string {
	return fmt.Sprintf("%s_rpc_%d", rb.cfg.QueuePrefix, rb.cfg.Node)
}

// 队列绑定key
func (rb *RabbitMQ) getRpcBindKey() string {
	return fmt.Sprintf("%s_rpc.%d.*", rb.cfg.QueuePrefix, rb.cfg.Node)
}

// 路由key
func (rb *RabbitMQ) getRoutingKey(node int) string {
	return fmt.Sprintf("%s.%d", rb.cfg.QueuePrefix, node)
}

// 路由key
func (rb *RabbitMQ) getRpcRoutingKey(node int) string {
	return fmt.Sprintf("%s_rpc.%d", rb.cfg.QueuePrefix, node)
}

// 异步发送消息
// 异步发送消息
func (rb *RabbitMQ) SendMessage(node int, body []byte) error {
	if rb.pubMessageChan == nil {
		return errors.New("not registered")
	}
	if rb.cfg.SendTimeout > 0 {
		select {
		case rb.pubMessageChan <- &request{
			node: node,
			body: body,
		}:
		case <-time.After(rb.cfg.SendTimeout):
			return ErrorTimeout
		}
	} else {
		rb.pubMessageChan <- &request{
			node: node,
			body: body,
		}
	}
	return nil
}

// rpc 服务调用
// 同步发送等待返回
func (rb *RabbitMQ) SyncMessage(node int, body []byte) ([]byte, error) {
	retChan := make(chan []byte, 1)
	rb.rpcRequestChan <- &request{
		node: node,
		body: body,
		ret:  retChan,
	}
	select {
	case b := <-retChan:
		if len(b) == len(timeoutByte) && string(b) == string(timeoutByte) {
			return nil, ErrorTimeout
		}
		return b, nil
	case <-time.After(rb.cfg.RPCTimeout):
		return nil, ErrorTimeout
	}
}

// 返回接收通道
func (rb *RabbitMQ) ReceiveMessage() []byte {
	return <-rb.subMessageChan
}

// 关闭
func (rb *RabbitMQ) Close() {
	if rb.cancel != nil {
		rb.cancel()
	}
}

// 重连
func redial(ctx context.Context, url, exchange, exchangeKind string) chan chan session {
	sessions := make(chan chan session)

	go func() {
		sess := make(chan session)
		defer close(sessions)

		for {
			select {
			case sessions <- sess:
			case <-ctx.Done():
				log.Infof("[rabbit] shutting down session factory")
				return
			}

			conn, err := amqp.Dial(url)
			if err != nil {
				log.Infof("[rabbit] cannot (re)dial: %v: %q", err, url)
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Infof("[rabbit] cannot create channel: %v", err)
			}

			if err := ch.ExchangeDeclare(exchange, exchangeKind, true, true, false, false, nil); err != nil {
				log.Infof("[rabbit] cannot declare %v exchange: %v", exchangeKind, err)
			}

			select {
			case sess <- session{conn, ch}:
			case <-ctx.Done():
				log.Infof("[rabbit] shutting down new session")
				return
			}
		}
	}()
	return sessions
}

// session
type session struct {
	*amqp.Connection
	*amqp.Channel
}

// 是否是成功的连接
func (s session) connected() bool {
	return s.Connection != nil && s.Channel != nil
}

// 连同连接和通道一起关闭
func (s session) close() {
	if s.Connection != nil {
		s.Connection.Close()
	}
}
