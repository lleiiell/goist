package goist

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestAwaitAll(t *testing.T) {
	inputs := []int{3, 4, 5}
	type ExampleResponse struct {
		data string
		err  error
	}

	responses := AwaitAll(inputs, func(input int) ExampleResponse {
		// Function can be sync or async as it is wrapped in a goroutine

		// Mock request to server each with different lengths
		randomTime := Rand().Intn(5)

		time.Sleep(time.Duration(randomTime) * time.Second)

		// Error cannot be thrown, but must be in struct
		var err error

		if randomTime > 1 {
			err = fmt.Errorf("Request was too slow")
		}

		return ExampleResponse{
			data: strconv.Itoa(input),
			err:  err,
		}
	})

	// [{3 <nil>} {4 0xc000484010} {5 0xc000484000}]
	fmt.Println(responses)
}
