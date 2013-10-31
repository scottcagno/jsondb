// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * dbgc.go :: database garbage collector
// *

package jsondb

import (
	"sync"
	"time"
)

type GC struct {
	gc map[string]int64
	ds *DataStore
	sync.Mutex
}

func InitGC(ds *DataStore) *GC {
	return &GC{
		gc: make(map[string]int64),
		ds: ds,
	}
}

func (self *GC) Run() {
	if len(self.gc) > 0 {
		self.Lock()
		for k, ttl := range self.gc {
			if ttl <= time.Now().Unix() {
				self.ds.del(&Request{Cmd: "del", Key: k, Val: nil})
				delete(self.gc, k)
			} else {
				break
			}
		}
		self.Unlock()
	}
	time.AfterFunc(time.Duration(1)*time.Second, func() { self.Run() })
}

func (self *GC) exp(k string, v int64) int64 {
	self.Lock()
	if _, ok := self.gc[k]; ok {
		self.gc[k] = v + time.Now().Unix()
		self.Unlock()
		return 1
	}
	self.Unlock()
	return -1
}

func (self *GC) ttl(k string) int64 {
	self.Lock()
	if ttl, ok := self.gc[k]; ok {
		self.Unlock()
		return ttl
	}
	self.Unlock()
	return -1
}
