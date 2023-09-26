package goist

import (
	"math/rand"
	"reflect"
	"sort"
	"time"
)

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

// SliceIntDiff2Right returns the elements in `a` that aren't in `b`.
func SliceIntDiff2Right(a, b []int) (c []int) {

	m := make(map[any]any)

	for _, v := range b {
		m[v] = 1
	}

	for _, v := range a {
		if _, ok := m[v]; !ok {
			c = append(c, v)
		}
	}

	return
}

func ToInterfaceSlice(sl interface{}) (isl []interface{}) {
	switch reflect.TypeOf(sl).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return
	}

	s := reflect.ValueOf(sl)
	for i := 0; i < s.Len(); i++ {
		isl = append(isl, s.Index(i).Interface())
	}

	return
}