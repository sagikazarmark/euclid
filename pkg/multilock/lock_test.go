package multilock

import (
	"runtime"
	"strconv"
	"testing"
)

func HammerLock(m *Lock, loops int, cdone chan bool) {
	for i := 0; i < loops; i++ {
		m.Lock("name")
		m.Unlock("name")
	}
	cdone <- true
}

func TestLock(t *testing.T) {
	if n := runtime.SetMutexProfileFraction(1); n != 0 {
		t.Logf("got mutexrate %d expected 0", n)
	}
	defer runtime.SetMutexProfileFraction(0)

	m := new(Lock)
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go HammerLock(m, 1000, c)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func HammerLockMultiple(m *Lock, cycle int, loops int, cdone chan bool) {
	for i := 0; i < loops; i++ {
		m.Lock(strconv.Itoa(cycle))
		m.Unlock(strconv.Itoa(cycle))
	}
	cdone <- true
}

func TestLock_CanHoldMultipleLocks(t *testing.T) {
	if n := runtime.SetMutexProfileFraction(1); n != 0 {
		t.Logf("got mutexrate %d expected 0", n)
	}
	defer runtime.SetMutexProfileFraction(0)

	m := new(Lock)
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go HammerLockMultiple(m, i, 1000, c)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}
