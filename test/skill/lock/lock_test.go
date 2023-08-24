package lock

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"
)

type GuardPolicy struct {
	mu      *sync.RWMutex
	errData map[*sql.DB]int
}

type sortPool struct {
	connPool gorm.ConnPool
	errCnt   int
}

func (s *GuardPolicy) countErr(db *gorm.DB) {
	if db.Error == nil || errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return
	}
	if ins, ok := db.Statement.ConnPool.(*sql.DB); ok {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.errData[ins] = s.errData[ins] + 1
	}
}

func (s *GuardPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	var x = make([]*sortPool, 0, len(connPools))
	for i := range connPools {
		p, ok := connPools[i].(*sql.DB)
		if !ok {
			x = append(x, &sortPool{connPool: connPools[i], errCnt: 0})
		} else {
			s.mu.RLock()
			defer s.mu.RUnlock()
			x = append(x, &sortPool{connPool: connPools[i], errCnt: s.errData[p]})
		}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].errCnt <= x[j].errCnt
	})
	return x[0].connPool
}

var mu *sync.RWMutex = &sync.RWMutex{}

func TestLock(t *testing.T) {

	lock()

	for i := 0; i < 100; i++ {
		go rlock(i)
	}

	go lock()
	go lock()
	go lock()
	go lock()
	go lock()
	time.Sleep(time.Second * 1)

	for i := 100; i < 200; i++ {
		go rlock(i)
	}

	/**ch := make(chan bool)
	<-ch*/

}

func lock() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("lock ...")
}

func rlock(i int) {
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(time.Second * 5)

	fmt.Printf("rlock %v ---\n", i)
}
