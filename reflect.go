package goist

import "reflect"

func IsNumber(x interface{}) (y bool) {

	switch reflect.TypeOf(x).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		y = true
	}

	return
}
