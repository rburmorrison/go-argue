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
