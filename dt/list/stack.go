package list

// Stack 栈
type Stack struct {
	l []interface{}
}

// NewStack 创建栈
func NewStack(cap int) *Stack {
	l := make([]interface{}, 0, cap)
	return &Stack{l}
}

// Destory 销毁
func (s *Stack) Destory() {
	s.l = nil
}

// Len 栈长
func (s *Stack) Len() int {
	return len(s.l)
}

// Top 返回栈顶
func (s *Stack) Top() interface{} {
	length := len(s.l)
	if length == 0 {
		return nil
	}
	return s.l[length-1]
}

// Push 进栈
func (s *Stack) Push(v interface{}) {
	s.l = append(s.l, v)
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	length := len(s.l)
	if length == 0 {
		return nil
	}
	ret := s.l[length-1]
	s.l = s.l[:length-1]
	return ret
}
