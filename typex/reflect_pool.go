package typex

import (
	"reflect"
	"sync"
)

var ReflectUsePool bool

// Reset defines Reset method for pooled object.
type ReflectReset interface {
	Reset()
}

var ReflectTypePools = &ReflectTypePool{
	pools: make(map[reflect.Type]*sync.Pool),
	New: func(t reflect.Type) interface{} {
		var argv reflect.Value

		if t.Kind() == reflect.Ptr { // reply must be ptr
			argv = reflect.New(t.Elem())
		} else {
			argv = reflect.New(t)
		}

		return argv.Interface()
	},
}

type ReflectTypePool struct {
	mu    sync.RWMutex
	pools map[reflect.Type]*sync.Pool
	New   func(t reflect.Type) interface{}
}

func (p *ReflectTypePool) Init(t reflect.Type) {
	tp := &sync.Pool{}
	tp.New = func() interface{} {
		return p.New(t)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pools[t] = tp
}

func (p *ReflectTypePool) Put(t reflect.Type, x interface{}) {
	if !ReflectUsePool {
		return
	}
	if o, ok := x.(ReflectReset); ok {
		o.Reset()
	}

	p.mu.RLock()
	pool := p.pools[t]
	p.mu.RUnlock()
	pool.Put(x)
}

func (p *ReflectTypePool) Get(t reflect.Type) interface{} {
	if !ReflectUsePool {
		return p.New(t)
	}
	p.mu.RLock()
	pool := p.pools[t]
	p.mu.RUnlock()

	return pool.Get()
}
