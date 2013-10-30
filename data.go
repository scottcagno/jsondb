// *
// * Copyright Scott Cagno, 2013. All rights reserved.
// * BSD Licensed. sites.google.com/site/bsdc3license
// *
// * data.go :: data store
// *

package jsondb

import (
	"sync"
)

type DataStore struct {
	ds map[string]interface{}
	sync.Mutex
}

func InitDataStore() *DataStore {
	self := &DataStore{
		ds: make(map[string]interface{}),
	}
	return self
}

func (self *DataStore) has(r *Request) Response {
	self.Lock()
	if _, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return Response{"code": 1, "val": "key found"}
	}
	self.Unlock()
	return Response{"code": -1, "val": "key not found"}
}

func (self *DataStore) add(r *Request) Response {
	self.Lock()
	if _, ok := self.ds[r.Key]; !ok {
		self.ds[r.Key] = r.Val
		self.Unlock()
		return Response{"code": 1, "val": "add ok"}
	}
	self.Unlock()
	return Response{"code": -1, "val": "add err"}
}

func (self *DataStore) set(r *Request) Response {
	self.Lock()
	self.ds[r.Key] = r.Val
	if _, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return Response{"code": 1, "val": "set ok"}
	}
	self.Unlock()
	return Response{"code": -1, "val": "set err"}
}

func (self *DataStore) get(r *Request) Response {
	self.Lock()
	if v, ok := self.ds[r.Key]; ok {
		self.Unlock()
		return Response{"code": 1, "val": v}
	}
	self.Unlock()
	return Response{"code": -1, "val": "get err"}
}

func (self *DataStore) del(r *Request) Response {
	self.Lock()
	delete(self.ds, r.Key)
	if _, ok := self.ds[r.Key]; !ok {
		self.Unlock()
		return Response{"code": 1, "val": "del ok"}
	}
	self.Unlock()
	return Response{"code": -1, "val": "del err"}
}

func (self *DataStore) err(r *Request) Response {
	return Response{"code": -1, "val": "unknown error"}
}
