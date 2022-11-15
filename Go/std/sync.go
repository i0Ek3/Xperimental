package std

import (
	"internal/race"
	"log"
	"sync/atomic"
	"unsafe"
)

// https://pkg.go.dev/sync@go1.19.3

type Locker interface {
	Lock()
	Unlock()
}

type noCopy struct{}
type notifyList struct{}
type copyChecker struct {
	uintptr
	copied bool
}

type Cond struct {
	noCopy  noCopy
	L       Locker
	notify  notifyList
	checker copyChecker
}

func (c *copyChecker) check() {
	if c.copied {
		panic("sync.Cond is copied")
	}
}

func runtime_notifyListAdd(l *notifyList) uint32     { return 0 }
func runtime_notifyListWait(l *notifyList, t uint32) {}
func runtime_notifyListNotifyAll(l *notifyList)      {}
func runtime_notifyListNotifyOne(l *notifyList)      {}

func NewCond(l Locker) *Cond {
	return &Cond{L: l}
}

func (c *Cond) Broadcast() {
	c.checker.check()
	runtime_notifyListNotifyAll(&c.notify)
}

func (c *Cond) Signal() {
	c.checker.check()
	runtime_notifyListNotifyOne(&c.notify)
}

func (c *Cond) Wait() {
	c.checker.check()
	t := runtime_notifyListAdd(&c.notify)
	c.L.Unlock()
	runtime_notifyListWait(&c.notify, t)
	c.L.Lock()
}

type Map struct {
	mu     Mutex
	read   atomic.Value // = readOnly
	dirty  map[any]*entry
	misses int
}

type readOnly struct {
	m       map[any]*entry
	amended bool
}

type entry struct {
	p unsafe.Pointer
}

func newEntry(i any) *entry {
	return &entry{p: unsafe.Pointer(&i)}
}

func (m *Map) Delete(key any) {
	m.LoadAndDelete(key)
}

func (m *Map) Load(key any) (value any, ok bool) {
	// first read
	read, _ := m.read.Load().(readOnly)
	e, ok := read.m[key]
	// not found
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly)
		e, ok = read.m[key]
		// not found
		if !ok && read.amended {
			// find in dirty
			e, ok = m.dirty[key]
			// update misses
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return nil, false
	}
	return e.load()
}

var expunged = new(any)

type Pointer struct {
	v unsafe.Pointer
}

// LoadPointer atomically loads *addr.
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) { return val }

func (e *entry) load() (value any, ok bool) {
	p := any(LoadPointer(&e.p))
	if p == nil || p == expunged {
		return nil, false
	}
	return p, true
}

func (m *Map) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(&readOnly{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

func (m *Map) Store(key, value any) {
	// fetch readOnly from m
	// read read.m[key], if ok and not delete, store key and val
	// otherwise Lock()
	// case1 loadReadOnly()
	// case2 m.dirty[key]
	// case3 else situations
	// Unlock()
}

func (m *Map) LoadAndDelete(key any) (value any, loaded bool) {
	// load key from readOnly and then delete the key
	return nil, false
}

func (m *Map) LoadOrStore(key, value any) (actual any, loaded any) {
	// if key exist then return, otherwise store and return
	// loaded = false means stored
	return actual, loaded
}

type Mutex struct {
	state int32
	sema  uint32
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

const rwmutexMaxReaders = 1 << 30

func (m *Mutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	m.lockSlow()
}

func (m *Mutex) lockSlow() {}

func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {}

func (m *Mutex) TryLock() bool {
	old := m.state
	if old&(mutexLocked|mutexStarving) != 0 {
		return false
	}
	if !atomic.CompareAndSwapInt32(&m.state, old, old|mutexLocked) {
		return false
	}
	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
	return true
}

type Once struct {
	done uint32
	m    Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

type Pool struct {
	noCopy     noCopy
	local      unsafe.Pointer
	localSize  uintptr
	victim     unsafe.Pointer
	victimSize uintptr
	New        func() any
}

func (p *Pool) Get() (x any) {
	if race.Enabled {
		race.Disable()
	}
	if x == nil {
		// x = pop from poolChain
		if x == nil {
			// _, pid := p.pin()
			// x = getSlow(pid)
		}
	}
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
	if x == nil && p.New != nil {
		x = p.New()
	}
	return x
}

// poolRaceAddr kinda complicated so we just return the result
func poolRaceAddr(x any) unsafe.Pointer {
	return unsafe.Pointer(&x)
}

func (p *Pool) Put(x any) {
	if x == nil {
		return
	}
	if race.Enabled {
		race.Disable()
	}
	// push key into poolChain
}

type RWMutex struct {
	w           Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount int32
	readerWait  int32
}

func runtime_SemacquireMutex(s *uint32, lifo bool, skipframes int) {}
func runtime_Semrelease(s *uint32, handoff bool, skipframes int)   {}

func (rw *RWMutex) Lock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	rw.w.Lock()
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_SemacquireMutex(&rw.writerSem, false, 0)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
}

func (rw *RWMutex) RLock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		runtime_SemacquireMutex(&rw.readerSem, false, 0)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
	}
}

func (rw *RWMutex) RLocker() Locker {
	return (*RWMutex)(rw)
}

func (rw *RWMutex) RUnlock() {
	if race.Enabled {
		_ = rw.w.state
		race.ReleaseMerge(unsafe.Pointer(&rw.writerSem))
		race.Disable()
	}
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// Outlined slow-path to allow the fast-path to be inlined
		rw.rUnlockSlow(r)
	}
	if race.Enabled {
		race.Enable()
	}
}

func (rw *RWMutex) rUnlockSlow(r int32) {}

func (rw *RWMutex) TryRLock() bool {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	for {
		c := atomic.LoadInt32(&rw.readerCount)
		if c < 0 {
			if race.Enabled {
				race.Enable()
			}
			return false
		}
		if atomic.CompareAndSwapInt32(&rw.readerCount, c, c+1) {
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(&rw.readerSem))
			}
			return true
		}
	}
}

func (rw *RWMutex) TryLock() bool {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	if !rw.w.TryLock() {
		if race.Enabled {
			race.Enable()
		}
		return false
	}
	if !atomic.CompareAndSwapInt32(&rw.readerCount, 0, -rwmutexMaxReaders) {
		rw.w.Unlock()
		if race.Enabled {
			race.Enable()
		}
		return false
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
	return true
}

func (rw *RWMutex) Unlock() {
	if race.Enabled {
		_ = rw.w.state
		race.Release(unsafe.Pointer(&rw.readerSem))
		race.Disable()
	}

	// Announce to readers there is no active writer.
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	if r >= rwmutexMaxReaders {
		race.Enable()
		log.Fatal("sync: Unlock of unlocked RWMutex")
	}
	// Unblock blocked readers, if any.
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false, 0)
	}
	// Allow other writers to proceed.
	rw.w.Unlock()
	if race.Enabled {
		race.Enable()
	}
}

type WaitGroup struct {
	noCopy noCopy
	state1 uint64
	state2 uint32
}

func (wg *WaitGroup) Add(delta int) {
	// fetch wg's state and then disable race
	// if state >> 32 < 0 or wrong waiters then panic
	// otherwise reset waiters count
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	// fetch wg's state and then disable race
	// for loop to deal with wait situation
	// if state >> 32 == 0 which means no need to wait
	// otherwise increment waiters count
}
