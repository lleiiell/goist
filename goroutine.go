package goist

import "sync"

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
