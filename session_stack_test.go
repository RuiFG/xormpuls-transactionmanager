package transaction_manager

import (
	. "github.com/xormplus/xorm"
	"testing"
)

func TestStackIsEmpty(t *testing.T) {
	stack := NewSessionStack()
	stack.Push(new(Session))
	if stack.IsEmpty() {
		t.Errorf("the size of stack should be one")
	}
	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("the size of stack should be zero")
	}
}
func TestStackPop(t *testing.T) {
	stack := NewSessionStack()
	for i := 0; i < 10; i++ {
		stack.Push(new(Session))
	}
	for i := 0; i < 10; i++ {
		stack.Pop()
	}
	if _, ok := stack.Pop(); ok {
		t.Errorf("stack pop function should return false")
	}
}
func TestStackEqual(t *testing.T) {
	stack := NewSessionStack()
	session := new(Session)
	stack.Push(session)
	sourceSession, _ := stack.Pop()
	if session != sourceSession {
		t.Errorf("session should be equal")
	}
}
func TestStackSize(t *testing.T) {
	stack := NewSessionStack()
	size := 10
	for i := 0; i < size; i++ {
		stack.Push(new(Session))
	}
	if stack.Size() != size {
		t.Errorf("the size of stack should be equal")
	}
}
func TestStackTop(t *testing.T) {
	stack := NewSessionStack()
	if _, ok := stack.Top(); ok {
		t.Errorf("top should return false")
	}
	for i := 0; i < 10; i++ {
		stack.Push(new(Session))
	}
	s := new(Session)
	stack.Push(s)
	if sourceSession, _ := stack.Top(); sourceSession != s {
		t.Errorf("session should be equal")
	}
}
func TestStackPush(t *testing.T) {
	stack := NewSessionStack()
	size := 20
	for i := 0; i < size; i++ {
		stack.Push(new(Session))
		if stack.Size() != i+1 {
			t.Errorf("push func error")
		}
	}
}
