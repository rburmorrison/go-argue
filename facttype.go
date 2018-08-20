package argue

import (
	"fmt"
	"reflect"
)

// FactType represents the data type that a fact will
// accept.
type FactType int

// FactType Values
const (
	FactTypeString = iota
	FactTypeBool
	FactTypeInt
	FactTypeInt64
	FactTypeUInt
	FactTypeUInt64
	FactTypeFloat32
	FactTypeFloat64
)

// GetFactType accepts an interface and will return
// it's equivelant FactType, or an error if it
// doesn't match a FactType.
func GetFactType(v interface{}) (FactType, error) {
	var t FactType
	ty := fmt.Sprintf("%s", reflect.TypeOf(v))
	switch ty {
	case "*string":
		t = FactTypeString
	case "*bool":
		t = FactTypeBool
	case "*int":
		t = FactTypeInt
	case "*int64":
		t = FactTypeInt64
	case "*uint":
		t = FactTypeUInt
	case "*uint64":
		t = FactTypeUInt64
	case "*float32":
		t = FactTypeFloat32
	case "*float64":
		t = FactTypeFloat64
	default:
		return FactType(-1), ErrInvalidType
	}

	return t, nil
}
