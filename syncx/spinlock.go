package syncx

import (
	"runtime"

	"go.uber.org/atomic"
)

type SpinLock interface {
	Lock()
	UnLock()
}

func NewSpinLock() SpinLock {
	return new(_SpinLock)
}

type _SpinLock struct {
	lock atomic.Uint32
}

func (l *_SpinLock) Lock() {
	for !l.TryLock() {
		runtime.Gosched()
	}
}

func (l *_SpinLock) TryLock() bool {
	return l.lock.CAS(0, 1)
}

func (l *_SpinLock) UnLock() {
	l.lock.Store(0)
}
