package goist

import (
	"reflect"
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
