package transaction_manager

import . "github.com/xormplus/xorm"

//return the session of the transaction manager
type DynamicSession func() *Session

//transaction function
type TransactionFunc func(transaction *Transaction) error

//routine's session stack interface
type SessionStack interface {
	Pop() (*Session, bool)
	Push(node *Session)
	Top() (*Session, bool)
	Size() int
	IsEmpty() bool
}

//session synchronize manager interface
type SynchronizeManager interface {
	Get() *Session
	Remove()
	Add(node *Session)
}

//transaction manager interface
type TransactionManager interface {
	SessionFunc() DynamicSession
	Do(do TransactionFunc, transactionDefinition ...int) error
}
