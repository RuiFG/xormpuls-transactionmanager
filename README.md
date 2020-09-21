# XormPlus-TransactionManager
######
xormplus transaction-manager is the xorm plus automatic transaction manager. With this library, 
you don’t need to worry about global transaction.

# Installation
```shell script
go get github.com/RuiFG/xormpuls-transactionmanager 
```
# Simple Example
```go
package main

import (
	"github.com/RuiFG/xormpuls-transactionmanager"
	"github.com/xormplus/xorm"
)

type Test struct {
	id          int
	description string
}

func main() {
	db, _ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	manager := transaction_manager.New(db)
	DynamicSession := manager.SessionFunc()
	test := Test{}
	_ = manager.Do(func(transaction *xorm.Transaction) error {
		return DynamicSession().Where("id = ?", 1).Find(&test)
	}, xorm.PROPAGATION_NEVER)
}
```

# License
This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.