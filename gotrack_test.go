package transaction_manager

import (
	"sync"
	"testing"
)

func TestCurGoroutineID(t *testing.T) {
	var group sync.WaitGroup
	group.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			if curGoroutineID() != curGoroutineID() {
				t.Errorf("cur go routine id should be equal")
			}
			group.Done()
		}()
	}
	group.Wait()
}
