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
