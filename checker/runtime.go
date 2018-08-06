package checker

import (
	"github.com/fabiorphp/htz"
	"runtime"
)

// Runtime returns the status of application runtime.
func Runtime(optional bool) htz.Checker {
	return func() *htz.Check {
		var mem runtime.MemStats

		runtime.ReadMemStats(&mem)

		details := map[string]interface{}{
			"goroutines":         runtime.NumGoroutine(),
			"memory_total_alloc": mem.TotalAlloc,
			"memory_heap_alloc":  mem.HeapAlloc,
		}

		return &htz.Check{
			Name:     "runtime",
			Type:     App,
			Status:   true,
			Optional: optional,
			Details:  details,
		}
	}
}
