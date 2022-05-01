package syncx

import "sync"

type SingleFlight interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
	DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
}

type _FlightGroup struct {
	calls map[string]*_Call
	lock  sync.Mutex
}

type _Call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

func NewSingleFlight() SingleFlight {
	return &_FlightGroup{
		calls: make(map[string]*_Call),
	}
}

func (g *_FlightGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	var (
		c  *_Call
		ok bool
	)
	g.lock.Lock()
	if c, ok = g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c = g.MakeCall(key, fn)
	return c.val, c.err
}

func (g *_FlightGroup) DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error) {
	var (
		c  *_Call
		ok bool
	)
	g.lock.Lock()
	if c, ok = g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c.val, false, c.err
	}

	c = g.MakeCall(key, fn)
	return c.val, true, c.err
}

func (g *_FlightGroup) MakeCall(key string, fn func() (interface{}, error)) *_Call {
	c := new(_Call)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	defer func() {
		g.lock.Lock()
		delete(g.calls, key)
		g.lock.Unlock()
		c.wg.Done()
	}()

	c.val, c.err = fn()
	return c
}
