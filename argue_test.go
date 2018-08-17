package argue

import "testing"

func TestBreakCammelCase(t *testing.T) {
	example := "BreakThisString"
	n := breakCammelCase(example)
	if n != "Break-This-String" {
		t.Errorf("BreakCammelCase was incorrect, got: %s, expected: %s", n, "Break-This-String")
	}
}
