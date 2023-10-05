package core

import (
	"sync"
)

type obarrayType struct {
	val map[string]*lispSymbol
	// TODO: Change for RWMutex
	lock *sync.Mutex
}

func newObarray(addLock bool) obarrayType {
	var lock *sync.Mutex
	if addLock {
		lock = &sync.Mutex{}
	}
	return obarrayType{
		val:  make(map[string]*lispSymbol),
		lock: lock,
	}
}

func (ob *obarrayType) loadOrStoreSymbol(sym *lispSymbol) (*lispSymbol, bool) {
	if ob.lock != nil {
		ob.lock.Lock()
		defer ob.lock.Unlock()
	}

	name := xSymbolName(sym)
	prev, existed := ob.val[name]
	if existed {
		return prev, true
	}

	ob.val[name] = sym
	return sym, false
}

func (ob *obarrayType) removeSymbol(name string) bool {
	if ob.lock != nil {
		ob.lock.Lock()
		defer ob.lock.Unlock()
	}

	_, existed := ob.val[name]
	if !existed {
		return false
	}

	delete(ob.val, name)
	return true
}

func (ob *obarrayType) containsSymbol(name string) bool {
	if ob.lock != nil {
		ob.lock.Lock()
		defer ob.lock.Unlock()
	}

	_, existed := ob.val[name]
	return existed
}
