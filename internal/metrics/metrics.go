package metrics

import "sync/atomic"

var requestCount int64

func IncRequests() {
    atomic.AddInt64(&requestCount, 1)
}

func GetRequests() int64 {
    return atomic.LoadInt64(&requestCount)
}
