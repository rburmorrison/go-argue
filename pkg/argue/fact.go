package argue

// FactType represents the data type that a rule will
// accept.
type FactType int

// FactType Values
const (
	FactTypeInt = iota
	FactTypeFloat
	FactTypeString
	FactTypeBool
)

// Fact represents a rule that an argument must
// follow when parsing command-line arguments. Facts
// are akin to flags.
type Fact struct {
	Type FactType
}
