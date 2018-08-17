package argue

import (
	"testing"
)

func TestSplitArguments(t *testing.T) {
	var tUInt uint
	var tInt int
	var tBool bool
	var tString = "Default text"
	var tPos string
	var tOther int

	agmt := NewArgument("This is a test of the argument library.", "2.0.0")
	agmt.AddFlagFact("uint", "this is a uint", &tUInt)
	agmt.AddFlagFact("int", "this is an integer", &tInt)
	agmt.AddFlagFact("bool", "this is a boolean", &tBool)
	agmt.AddFlagFact("string", "this is a string", &tString)
	agmt.AddPositionalFact("pos", "this is a positional string", &tPos).SetRequired(true)
	agmt.AddPositionalFact("other", "this is another int", &tOther)

	arguments := []string{"--int", "123", "--string", "test string", "asdf", "--other"}
	ps, fm := agmt.SplitArguments(arguments)
	if len(fm) != 3 {
		t.Errorf("SplitArguments was incorrect, got: %d flags, expected: %d flags", len(fm), 3)
	}

	if len(ps) != 1 {
		t.Errorf("SplitArguments was incorrect, got: %d positionals, expected %d positionals", len(ps), 1)
	}

	if fm["--string"] != "test string" {
		t.Errorf("SplitArguments was incorrect: got: fm[\"--string\"] != \"test string\", expected: fm[\"--string\"] == \"test string\"")
	}
}

func TestDisputeCustom(t *testing.T) {
	var tUInt uint
	var tInt int
	var tBool bool
	var tString = "Default text"
	var tPos string
	var tOther int

	agmt := NewArgument("This is a test of the argument library.", "2.0.0")
	agmt.AddFlagFact("uint", "this is a uint", &tUInt)
	agmt.AddFlagFact("int", "this is an integer", &tInt)
	agmt.AddFlagFact("bool", "this is a boolean", &tBool)
	agmt.AddFlagFact("string", "this is a string", &tString)
	agmt.AddPositionalFact("pos", "this is a positional string", &tPos).SetRequired(true)
	agmt.AddPositionalFact("other", "this is another int", &tOther)

	arguments := []string{"--int", "123", "--string", "test string", "asdf"}
	err := agmt.DisputeCustom(arguments, false)
	if err != ErrMissingPositionals {
		t.Errorf("DisputeCustom was incorrect, expected: ErrMissingPositionals, got %v", err)
	}

	arguments = []string{"--int", "123", "--string", "test string", "asdf", "123", "--other"}
	err = agmt.DisputeCustom(arguments, false)
	if err != ErrUnknownFlag {
		t.Errorf("DisputeCustom was incorrect, expected: ErrUnknownFlag, got %v", err)
	}
}
