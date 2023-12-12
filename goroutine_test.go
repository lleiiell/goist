package goist

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

/*
=== RUN   TestConcurrencyChan
2023-12-12 14:23:52.49485124 +0800 CST m=+0.006995545 10
2023-12-12 14:23:52.494867824 +0800 CST m=+0.007012157 10
2023-12-12 14:23:52.494954694 +0800 CST m=+0.007099011 10
2023-12-12 14:23:52.995799262 +0800 CST m=+0.507943579 10
2023-12-12 14:23:52.995826921 +0800 CST m=+0.507971255 10
2023-12-12 14:23:52.995746334 +0800 CST m=+0.507890673 10
2023-12-12 14:23:53.496574924 +0800 CST m=+1.008719263 10
2023-12-12 14:23:53.49662873 +0800 CST m=+1.008773045 10
2023-12-12 14:23:53.496621925 +0800 CST m=+1.008766249 10
2023-12-12 14:23:53.997333913 +0800 CST m=+1.509478254 10
--- PASS: TestConcurrencyChan (2.00s)
*/
func TestConcurrencyChan(t *testing.T) {
	ch := make(chan int, 3)

	sc := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		sc.Add(1)
		go func() {
			ch <- 1
			defer func() {
				<-ch
				sc.Done()
			}()

			fmt.Println(time.Now(), i)
			time.Sleep(time.Millisecond * 500)
		}()
	}

	sc.Wait()
}

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
func TestConcurrencyTick(t *testing.T) {
	for i := 0; i < 10; i++ {
		<-concurrencyTick(2)
		fmt.Println(time.Now(), i)
	}
}

/*
=== RUN   TestAwaitAllWRL
2023-12-12 14:40:47.881207186 +0800 CST m=+0.005446728 9
2023-12-12 14:40:47.88120844 +0800 CST m=+0.005448021 4
2023-12-12 14:40:48.081943066 +0800 CST m=+0.206182640 1
2023-12-12 14:40:48.081897088 +0800 CST m=+0.206136643 0
2023-12-12 14:40:48.282449929 +0800 CST m=+0.406689488 2
2023-12-12 14:40:48.282497664 +0800 CST m=+0.406737230 3
2023-12-12 14:40:48.482936089 +0800 CST m=+0.607175652 7
2023-12-12 14:40:48.482974548 +0800 CST m=+0.607214139 8
2023-12-12 14:40:48.683342967 +0800 CST m=+0.807582522 5
2023-12-12 14:40:48.683410421 +0800 CST m=+0.807649992 6
2023-12-12 14:40:48.883885259 +0800 CST m=+1.008124824 [{0 <nil>} {1 <nil>} {2 <nil>} {3 <nil>} {4 <nil>} {5 <nil>} {6 <nil>} {7 <nil>} {8 <nil>} {9 <nil>}]
--- PASS: TestAwaitAllWRL (1.00s)
*/
func TestAwaitAllWRL(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	type ExampleResponse struct {
		data string
		err  error
	}

	responses := AwaitAllWRL(inputs, func(input int) ExampleResponse {

		// Error cannot be thrown, but must be in struct
		var err error

		fmt.Println(time.Now(), input)

		time.Sleep(200 * time.Millisecond)

		return ExampleResponse{
			data: strconv.Itoa(input),
			err:  err,
		}
	}, 2)

	fmt.Println(time.Now(), responses)
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
