package metrics

import (
	"sync/atomic"
)

var (
	brasilAPICount uint64
	viaCEpCount    uint64
)

func IncrementBrasilAPI() {
	atomic.AddUint64(&brasilAPICount, 1)
}

func IncrementViaCEP() {
	atomic.AddUint64(&viaCEpCount, 1)
}

func GetBrasilAPI() uint64 {
	return atomic.LoadUint64(&brasilAPICount)
}

func GetViaCEP() uint64 {
	return atomic.LoadUint64(&viaCEpCount)
}
