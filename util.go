package goist

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"time"
)

func RateLimiter(limit int, duration time.Duration) func() bool {
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

func Unique(slicePtr interface{}) {

	var j []int
	m := make(map[interface{}]bool)

	if reflect.TypeOf(slicePtr).Elem().Kind() != reflect.Slice {
		return
	}

	s := reflect.ValueOf(slicePtr).Elem()
	if s.Len() < 1 {
		return
	}

	for i := 0; i < s.Len(); i++ {
		if _, ok := m[s.Index(i).Interface()]; ok {
			j = append(j, i)
			continue
		}

		m[s.Index(i).Interface()] = true

	}

	sort.Slice(j, func(i, j int) bool {
		return i > j
	})

	for _, v := range j {
		s = reflect.AppendSlice(s.Slice(0, v), s.Slice(v+1, s.Len()))
	}

	reflect.ValueOf(slicePtr).Elem().Set(s)

}

func Contains(slice interface{}, item interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

func Rand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Retry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(Rand().Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}

	return nil
}

func retryWithBackoff(maxRetries int, backoffFactor time.Duration, operation func() error) error {
	for i := 0; i < maxRetries; i++ {
		err := operation()
		if err == nil {
			return nil
		}

		// calc backoff duration
		backoffTime := backoffFactor * time.Duration(1<<i) // exponential backoff
		fmt.Printf("Attempt %d failed: %s. Retrying in %s...\n", i+1, err, backoffTime)

		time.Sleep(backoffTime + time.Duration(rand.Int63n(int64(backoffFactor))))
	}
	return errors.New("max retries exceeded")
}
