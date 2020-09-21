package transaction_manager

import (
	"fmt"
	. "github.com/xormplus/xorm"
	"sync"
)

type synchronizeManager struct {
	syncMap sync.Map
	engine  EngineInterface
}

func (s *synchronizeManager) curRoutineStack() SessionStack {
	actual, _ := s.syncMap.LoadOrStore(curGoroutineID(), NewSessionStack())
	stack := actual.(SessionStack)
	return stack
}

func (s *synchronizeManager) Get() *Session {
	stack := s.curRoutineStack()
	if stack.IsEmpty() {
		return s.engine.NewSession()
	}

	if curSession, ok := stack.Top(); !ok {
		panic(fmt.Errorf("get session error,sessionStack size is %d", stack.Size()))
	} else {
		return curSession
	}
}

func (s *synchronizeManager) Remove() {
	stack := s.curRoutineStack()
	if _, ok := stack.Pop(); !ok {
		panic("the size of stack is zero")
	}
	if stack.IsEmpty() {
		s.syncMap.Delete(curGoroutineID())
	}
}

func (s *synchronizeManager) Add(node *Session) {
	stack := s.curRoutineStack()
	stack.Push(node)
}

func NewSynchronizeManager(engineInterface EngineInterface) SynchronizeManager {
	return &synchronizeManager{engine: engineInterface}
}
