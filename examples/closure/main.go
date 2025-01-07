package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	println("rateLimiter")
	simulateReqs()

	println("adder")
	adder()

	println("counter")
	counter()
}

func simulateReqs() {
	rateLimiter := rateLimiter(3, 10*time.Second)

	// Simulate requests
	for i := 1; i <= 10; i++ {
		if rateLimiter() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate limit exceeded\n", i)
		}
		time.Sleep(2 * time.Second) // Simulate time between requests
	}

}

func rateLimiter(limit int, duration time.Duration) func() bool {
	var mu sync.Mutex
	var timestamps []time.Time

	return func() bool {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		// Remove timestamps older than the specified duration
		validUntil := now.Add(-duration)
		for len(timestamps) > 0 && timestamps[0].Before(validUntil) {
			timestamps = timestamps[1:] // Remove old timestamps
		}

		// Check if we can add a new request
		if len(timestamps) < limit {
			timestamps = append(timestamps, now)
			return true // Allowed
		}

		return false // Rate limit exceeded
	}
}

/*
2
4
6
4
6
8
*/
func adder() {
	add := func(x int) func(y int) int {
		return func(y int) int {
			return x + y
		}
	}

	add1 := add(1)
	println(add1(1))
	println(add1(3))
	println(add1(5))

	add2 := add(2)
	println(add2(2))
	println(add2(4))
	println(add2(6))
}

/*
1
2
3
1
2
*/
func counter() {
	count := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	incr1 := count()

	fmt.Println(incr1())
	fmt.Println(incr1())
	fmt.Println(incr1())

	incr2 := count()
	fmt.Println(incr2())
	fmt.Println(incr2())
}
