package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	// fn("you must face the diffcullty")
	// field := val.Field(0)
	// fn(field.String())
	val := getValue(x)

	// 	if field.Kind() == reflect.String {
	// 		fn(field.String())
	// 	}
	// 	if field.Kind() == reflect.Struct {
	// 		walk(field.Interface(), fn)
	// 	}
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		fmt.Printf("%v ==== %T\n", val, val.Type())
		fmt.Println("====================")
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			fmt.Printf("%v, %t\n", v, ok)
			walkValue(v)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
