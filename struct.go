package goist

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Struct2Map(in interface{}) (out map[string]interface{}, err error) {

	out = map[string]interface{}{}

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		err = fmt.Errorf("only accepts struct, got %T", v)
		return
	}

	tp := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := tp.Field(i)

		if tv := fi.Tag.Get("json"); tv != "" {

			tv, _ := strings.CutSuffix(tv, ",omitempty")

			out[tv] = v.Field(i).Interface()

		}
	}

	return
}

func Struct2MapMarshal(st interface{}) (m map[string]interface{}, err error) {

	bt, err := json.Marshal(st)
	if err != nil {
		return
	}

	err = json.Unmarshal(bt, &m)

	return
}

func Struct2Str(st interface{}, tag, sep string) (s string, err error) {
	v := reflect.ValueOf(st)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		err = fmt.Errorf("only accepts struct, got %T", v)
		return
	}

	if tag == "" {
		tag = "json"
	}

	if sep == "" {
		sep = ","
	}

	tp := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := tp.Field(i)

		if tv := fi.Tag.Get(tag); tv != "" {

			tv, _ := strings.CutSuffix(tv, ",omitempty")

			s += fmt.Sprintf("%s,", tv)

		} else {
			s += fmt.Sprintf("%s,", fi.Name)
		}
	}

	s = strings.TrimRight(s, sep)
	return
}
