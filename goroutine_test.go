package goist

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
2023-12-12 12:28:28.027853034 +0800 CST m=+0.505727705 0
2023-12-12 12:28:28.528673829 +0800 CST m=+1.006548499 1
2023-12-12 12:28:29.029433145 +0800 CST m=+1.507307819 2
2023-12-12 12:28:29.529954763 +0800 CST m=+2.007829440 3
2023-12-12 12:28:30.030517426 +0800 CST m=+2.508392097 4
2023-12-12 12:28:30.531494473 +0800 CST m=+3.009369206 5
2023-12-12 12:28:31.032528187 +0800 CST m=+3.510402858 6
2023-12-12 12:28:31.533658384 +0800 CST m=+4.011533047 7
2023-12-12 12:28:32.033845974 +0800 CST m=+4.511720645 8
2023-12-12 12:28:32.534253095 +0800 CST m=+5.012127766 9
*/
func TestSimpleTickConcurrency(t *testing.T) {
	for i := 0; i < 10; i++ {
		<-time.Tick(500 * time.Millisecond)
		fmt.Println(time.Now(), i)
	}
}

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
