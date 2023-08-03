package lock

import (
	"database/sql"
	"errors"
	"sort"
	"sync"

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
			errCnt := s.errData[p]
			s.mu.RUnlock()
			x = append(x, &sortPool{connPool: connPools[i], errCnt: errCnt})
		}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].errCnt <= x[j].errCnt
	})
	return x[0].connPool
}
