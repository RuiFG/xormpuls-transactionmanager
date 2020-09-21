package transaction_manager

import . "github.com/xormplus/xorm"

type sessionStack struct {
	items []*Session
}

func (s *sessionStack) Push(data *Session) {
	s.items = append(s.items, data)
}
func (s *sessionStack) Pop() (*Session, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	data := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return data, true
}

func (s sessionStack) Size() int {
	return len(s.items)
}

func (s sessionStack) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s sessionStack) Top() (*Session, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

func NewSessionStack() SessionStack {
	return &sessionStack{}
}
