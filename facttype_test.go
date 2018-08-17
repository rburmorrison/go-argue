package argue

import "testing"

func TestGetFactType(t *testing.T) {
	var s string
	ft, _ := GetFactType(&s)
	if ft != FactTypeString {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeString)
	}

	var b bool
	ft, _ = GetFactType(&b)
	if ft != FactTypeBool {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeBool)
	}

	var i int
	ft, _ = GetFactType(&i)
	if ft != FactTypeInt {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeInt)
	}

	var i64 int64
	ft, _ = GetFactType(&i64)
	if ft != FactTypeInt64 {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeInt64)
	}

	var u uint
	ft, _ = GetFactType(&u)
	if ft != FactTypeUInt {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeUInt)
	}

	var u64 uint64
	ft, _ = GetFactType(&u64)
	if ft != FactTypeUInt64 {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeUInt64)
	}

	var f32 float32
	ft, _ = GetFactType(&f32)
	if ft != FactTypeFloat32 {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeFloat32)
	}

	var f64 float64
	ft, _ = GetFactType(&f64)
	if ft != FactTypeFloat64 {
		t.Errorf("GetFactType was incorrect, got: %v, expected: %v", ft, FactTypeFloat64)
	}
}
