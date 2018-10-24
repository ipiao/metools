package notify

import (
	"encoding/json"
	"net/smtp"
	"time"

	"github.com/ipiao/metools/mencode"
	"github.com/jordan-wright/email"
)

var (
	notifyqueue = make(chan Message)
	epool       *email.Pool
)

func init() {
	pwd := mencode.Base64Decode("eWtrQCMxMDAx")
	epool = email.NewPool("smtp.qq.com:587", 5, smtp.PlainAuth("", "530151330@qq.com", pwd, "smtp.qq.com"))
	go handleNotify()
}

// Message for message
type Message struct {
	Error     string `json:"error"`
	OpContent string `json:"op_content"`
}

// Notify notify
func (m Message) Notify() {
	notifyqueue <- m
}

// Notify notify
func Notify(content string, err error) {
	var errr string
	if err != nil {
		errr = err.Error()
	}
	msg := &Message{Error: errr, OpContent: content}
	msg.Notify()
}

func handleNotify() {
	for {
		select {
		case msg := <-notifyqueue:
			notifyEmail(msg)
		}
	}
}

func notifyEmail(m Message) {
	e := email.NewEmail()
	e.Headers.Set("Content-type", "text/plain;charset=UTF-8")
	e.From = "Yu Kaokao <530151330@qq.com>"
	e.To = []string{"yukk046@zhixubao.com"}
	e.Subject = "notify from my programer"
	e.Text, _ = json.Marshal(&m)
	epool.Send(e, time.Second*10)
}
