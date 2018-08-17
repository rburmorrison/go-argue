// Package mirror implements functions to make using
// reflection easier.
package mirror

import "reflect"

// IsPointer accepts and interface and will return
// true if it is a pointer. False otherwise.
func IsPointer(v interface{}) bool {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return false
	}

	return true
}

// IsPointerToStruct accepts an interface and will
// return true if it is a pointer to a struct.
func IsPointerToStruct(v interface{}) bool {
	if reflect.Indirect(reflect.ValueOf(v)).Kind() != reflect.Struct {
		return false
	}

	return true
}
