package goist

import (
	"sync"
	"time"
)

func AwaitAll[Input any, Response any](inputs []Input, asyncFunc func(Input) Response) []Response {
	responses := make([]Response, len(inputs))

	var waitGroup sync.WaitGroup

	for i, input := range inputs {
		waitGroup.Add(1)
		go func(i int, input Input) {
			responses[i] = asyncFunc(input)
			waitGroup.Done()
		}(i, input)
	}

	waitGroup.Wait()

	return responses
}

// AwaitAllWRL AwaitAll with Rate Limit
func AwaitAllWRL[Input any, Response any](inputs []Input, asyncFunc func(Input) Response, concurrency int) []Response {
	responses := make([]Response, len(inputs))
	if concurrency < 1 {
		concurrency = 1
	}
	cch := make(chan int, concurrency)

	var waitGroup sync.WaitGroup

	for i, input := range inputs {
		waitGroup.Add(1)
		go func(i int, input Input) {
			cch <- 1
			defer func() {
				waitGroup.Done()
				<-cch
			}()
			responses[i] = asyncFunc(input)

		}(i, input)
	}

	waitGroup.Wait()

	return responses
}

func concurrencyTick(n int) <-chan time.Time {
	if n < 1 {
		return time.Tick(1 * time.Second)
	}
	c := 1e9 / n
	return time.Tick(time.Nanosecond * time.Duration(c))
}
