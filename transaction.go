package transaction_manager

import (
	"fmt"
	. "github.com/xormplus/xorm"
)

type transactionManager struct {
	manager SynchronizeManager
	engine  EngineInterface
}

func (s transactionManager) SessionFunc() DynamicSession {
	return func() *Session {
		return s.manager.Get()
	}
}

func (s transactionManager) Do(do TransactionFunc, transactionDefinition ...int) error {
	var result error
	var session *Session
	//remove session on sessionStack
	defer s.manager.Remove()
	// close session
	defer func() {
		if err := session.Close(); err != nil {
			panic(fmt.Errorf("close transaction error:%v", err))
		}
	}()
	// rollback or nothing to do when catch err
	defer func() {
		if err := recover(); err != nil {
			result = err.(error)
			if rollbackErr := session.Rollback(); rollbackErr != nil {
				panic(fmt.Errorf("rollback transaction error:%v", err))
			}
		}
	}()
	var sourceSession *Session
	sourceSession = s.manager.Get()
	trans, err := sourceSession.
		BeginTrans(transactionDefinition...)
	if err != nil {
		panic(fmt.Errorf("begin transacation error"))
	}
	session = trans.Session()
	s.manager.Add(session)
	if err := do(trans); err != nil {
		panic(err)
	}
	if err := trans.CommitTrans(); err != nil {
		panic(err)
	}
	return result
}

func New(param EngineInterface) TransactionManager {
	return transactionManager{manager: NewSynchronizeManager(param), engine: param}
}
