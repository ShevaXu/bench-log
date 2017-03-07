package log_test

import (
	"bytes"
	"log"
	"testing"
	"sync"

	newLog "github.com/ShevaXu/bench-log"
)

// TODO -> $ go test -bench="."


// Benchmark mutex for reference
func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

// Comparison of std log - no runtime.Caller()

func BenchmarkLog(b *testing.B) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.LstdFlags)
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

func BenchmarkNewLog(b *testing.B) {
	var buf bytes.Buffer
	logger := newLog.New(&buf, "", log.LstdFlags)
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

// Comparison of single goroutine logging

func BenchmarkLogWithFileLine(b *testing.B) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lshortfile)
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

func BenchmarkNewLogWithFileLine(b *testing.B) {
	var buf bytes.Buffer
	logger := newLog.New(&buf, "", log.Lshortfile)
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

// Comparison of multiple goroutines logging

const M = 5

func BenchmarkLogWithFileLineMulti(b *testing.B) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lshortfile)
	for i := 0; i < M; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				logger.Println("hello")
			}
		}()
	}
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

func BenchmarkNewLogWithFileLineMulti(b *testing.B) {
	var buf bytes.Buffer
	logger := newLog.New(&buf, "", log.Lshortfile)
	for i := 0; i < M; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				logger.Println("hello")
			}
		}()
	}
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

// Comparison of multiple goroutines logging - more goroutines

func BenchmarkLogWithFileLineMulti2(b *testing.B) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lshortfile)
	for i := 0; i < M * 2; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				logger.Println("hello")
			}
		}()
	}
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}

func BenchmarkNewLogWithFileLineMulti2(b *testing.B) {
	var buf bytes.Buffer
	logger := newLog.New(&buf, "", log.Lshortfile)
	for i := 0; i < M * 2; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				logger.Println("hello")
			}
		}()
	}
	for i := 0; i < b.N; i++ {
		logger.Println("hello")
	}
}
