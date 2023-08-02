package goist

import (
	"math/rand"
	"reflect"
	"time"
)

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
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rn
}
