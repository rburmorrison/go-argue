package argue

import "testing"

func TestStandardizeFactName(t *testing.T) {
	testS := "A Test   ----------     Name"
	expectedS := "a-test-name"
	s := StandardizeFactName(testS)
	if s != expectedS {
		t.Errorf("StandardizeFactName was incorrect, got: %s, want: %s", testS, expectedS)
	}

	testS = "A Test   ----------     Name"
	expectedS = "ATESTNAME"
	s = UpperFactName(testS)
	if s != expectedS {
		t.Errorf("UpperFactName was incorrect, got: %s, want: %s", testS, expectedS)
	}
}
