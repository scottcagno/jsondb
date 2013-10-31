// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * data.go :: data store
// *

package jsondb

import (
	"reflect"
	"sync"
)

type Request struct {
	Cmd string      `json:"cmd"`
	Key string      `json:"key"`
	Val interface{} `json:"val"`
}

type DataStore struct {
	ds map[string]interface{}
	gc *GC
	sync.Mutex
}

func InitDataStore() *DataStore {
	self := &DataStore{
		ds: make(map[string]interface{}),
	}
	self.gc = InitGC(self)
	go self.gc.Run()
	return self
}

func (self *DataStore) exp(r *Request) interface{} {
	var ret int64 = -1
	if reflect.TypeOf(r.Val).Kind() == reflect.Float64 {
		ret = self.gc.exp(r.Key, int64(r.Val.(float64)))
	}
	return ret
}

func (self *DataStore) ttl(r *Request) interface{} {
	return self.gc.ttl(r.Key)
}

func (self *DataStore) has(r *Request) interface{} {
	self.Lock()
	if _, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return 1
	}
	self.Unlock()
	return -1
}

func (self *DataStore) add(r *Request) interface{} {
	self.Lock()
	if _, ok := self.ds[r.Key]; !ok {
		self.ds[r.Key] = r.Val
		self.Unlock()
		return 1
	}
	self.Unlock()
	return -1
}

func (self *DataStore) set(r *Request) interface{} {
	self.Lock()
	self.ds[r.Key] = r.Val
	if _, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return 1
	}
	self.Unlock()
	return -1
}

func (self *DataStore) get(r *Request) interface{} {
	self.Lock()
	if v, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return v
	}
	self.Unlock()
	return -1
}

func (self *DataStore) del(r *Request) interface{} {
	self.Lock()
	delete(self.ds, r.Key)
	if _, ok := self.ds[r.Key]; !ok {
		self.Unlock()
		return 1
	}
	self.Unlock()
	return -1
}

func (self *DataStore) err(r *Request) interface{} {
	return -1
}
