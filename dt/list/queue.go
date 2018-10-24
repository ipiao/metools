package list

// Queue 队列
type Queue struct {
	l []interface{}
}

// NewQueue 创建队列
func NewQueue(cap int) *Queue {
	l := make([]interface{}, 0, cap)
	return &Queue{l}
}

// Destory 销毁
func (q *Queue) Destory() {
	q.l = nil
}

// Len 栈长
func (q *Queue) Len() int {
	return len(q.l)
}

// Head 返回堆头
func (q *Queue) Head() interface{} {
	length := len(q.l)
	if length == 0 {
		return nil
	}
	return q.l[0]
}

// Enter 进队
func (q *Queue) Enter(v interface{}) {
	q.l = append(q.l, v)
}

// Exit 出堆
func (q *Queue) Exit() interface{} {
	length := len(q.l)
	if length == 0 {
		return nil
	}
	ret := q.l[length-1]
	q.l = q.l[:length-1]
	return ret
}
