package std

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type common struct {
	mu       sync.Mutex
	output   []byte    // Output generated by test or benchmark.
	w        io.Writer // For flushToParent.
	failed   bool      // Test or benchmark has failed.
	skipped  bool      // Test or benchmark has been skipped.
	done     bool      // Test is finished and all subtests have completed.
	finished bool      // Test function has completed.
	inFuzzFn bool      // Whether the fuzz target, if this is one, is running.

	bench      bool        // Whether the current test is a benchmark.
	hasSub     atomic.Bool // whether there are sub-benchmarks.
	runner     string      // Function name of tRunner running the test.
	isParallel bool        // Whether the test is parallel.

	parent *common
	name   string // Name of test or benchmark.
}

type T struct {
	common
	isParallel bool
	isEnvSet   bool
	context    *testContext
}

// testContext holds all fields that are common to all tests.
type testContext struct {
	match    *matcher
	deadline time.Time

	// isFuzzing is true in the context used when generating random inputs
	// for fuzz targets. isFuzzing is false when running normal tests and
	// when running fuzz tests as unit tests (without -fuzz or when -fuzz
	// does not match).
	isFuzzing bool

	mu sync.Mutex

	// Channel used to signal tests that are ready to be run in parallel.
	startParallel chan bool

	// running is the number of tests currently running in parallel.
	running int

	// numWaiting is the number tests waiting to be run in parallel.
	numWaiting int

	// maxParallel is a copy of the parallel flag.
	maxParallel int
}

type matcher struct {
	filter    filterMatch
	skip      filterMatch
	matchFunc func(pat, str string) (bool, error)

	mu sync.Mutex

	// subNames is used to deduplicate subtest names.
	subNames map[string]int32
}

type filterMatch interface {
	matches(name []string, matchString func(pat, str string) (bool, error)) (ok, partial bool)
	verify(name string, matchString func(pat, str string) (bool, error)) error
}

type B struct {
	common
	importPath       string // import path of the package containing the benchmark
	context          *benchContext
	N                int
	previousN        int           // number of iterations in the previous run
	previousDuration time.Duration // total duration of the previous run
	benchFunc        func(b *B)
	benchTime        durationOrCountFlag
	bytes            int64
	missingBytes     bool // one of the subbenchmarks does not have bytes set.
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int // RunParallel creates parallelism*GOMAXPROCS goroutines
	// The initial states of memStats.Mallocs and memStats.TotalAlloc.
	startAllocs uint64
	startBytes  uint64
	// The net total of this test after being run.
	netAllocs uint64
	netBytes  uint64
	// Extra metrics collected by ReportMetric.
	extra map[string]float64
}

type benchContext struct {
	match *matcher

	maxLen int // The largest recorded benchmark name.
	extLen int // Maximum extension length.
}

type durationOrCountFlag struct {
	d         time.Duration
	n         int
	allowZero bool
}

type BenchmarkResult struct {
	N         int           // The number of iterations.
	T         time.Duration // The total time taken.
	Bytes     int64         // Bytes processed in one iteration.
	MemAllocs uint64        // The total number of memory allocations.
	MemBytes  uint64        // The total number of bytes allocated.

	// Extra records additional metrics reported by ReportMetric.
	Extra map[string]float64
}

func (c *common) checkFuzzFn(name string) {
	if c.inFuzzFn {
		panic(fmt.Sprintf("testing: f.%s was called inside the fuzz target, use t.%s instead", name, name))
	}
}

func (c *common) Errorf(format string, args ...any) {
	c.checkFuzzFn("Errorf")
	c.log(fmt.Sprintf(format, args...))
	c.Fail()
}

func (c *common) Fatalf(format string, args ...any) {
	c.checkFuzzFn("Fatalf")
	c.log(fmt.Sprintf(format, args...))
	c.FailNow()
}

func (c *common) Skipf(format string, args ...any) {
	c.checkFuzzFn("Skipf")
	c.log(fmt.Sprintf(format, args...))
	c.SkipNow()
}

func (c *common) SkipNow() {
	c.checkFuzzFn("SkipNow")
	c.mu.Lock()
	c.skipped = true
	c.finished = true
	c.mu.Unlock()
	runtime.Goexit()
}

func (c *common) FailNow() {
	c.checkFuzzFn("FailNow")
	c.Fail()
	c.mu.Lock()
	c.finished = true
	c.mu.Unlock()
	runtime.Goexit()
}

// log generates the output. It's always at the same stack depth 3.
func (c *common) log(s string) {}

func (c *common) Fail() {
	if c.parent != nil {
		c.parent.Fail()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// c.done needs to be locked to synchronize checks to c.done in parent tests.
	if c.done {
		panic("Fail in goroutine after " + c.name + " has completed")
	}
	c.failed = true
}