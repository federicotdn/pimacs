package core

import (
	"sync"
)

type obarrayType struct {
	val  map[string]*lispSymbol
	lock *sync.Mutex
}

func newObarray(lock bool) obarrayType {
	return obarrayType{
		val:  make(map[string]*lispSymbol),
		lock: &sync.Mutex{},
	}
}

func (ob *obarrayType) loadOrStoreSymbol(sym *lispSymbol) (*lispSymbol, bool) {
	if ob.lock != nil {
		ob.lock.Lock()
		defer ob.lock.Unlock()
	}

	prev, existed := ob.val[sym.name]
	if existed {
		return prev, true
	}

	ob.val[sym.name] = sym
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
