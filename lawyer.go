package argue

import (
	"os"
	"strings"
)

// Lawyer represents an entity that can parse through
// sub-commands. A Layer is essentially a manager for
// multiple arguments.
type Lawyer struct {
	Description  string
	Version      string
	SubArguments []SubArgument
	ShowDesc     bool
	ShowVersion  bool

	facts []Fact
}

// NewLawyer returns a new Lawyer with the version
// and description provided.
func NewLawyer(desc string, version string) Lawyer {
	var law Lawyer
	law.Description = desc
	law.Version = version
	law.ShowDesc = true
	law.ShowVersion = true
	return law
}

// NewEmptyLawyer returns a new Lawyer without a
// description or version to display.
func NewEmptyLawyer() Lawyer {
	law := NewLawyer("", "")
	law.ShowDesc = false
	law.ShowVersion = false
	return law
}

// AddArgument offers a new argument to the Lawyer
// with the passed parameters: name, help, and the
// argument to add.
func (l *Lawyer) AddArgument(n string, h string, arg Argument) *SubArgument {
	if l.NameExists(n) {
		panic("argue: this name already exists as a sub-argument")
	}

	var sarg SubArgument
	sarg.Name = n
	sarg.Help = h
	sarg.Argument = arg

	l.SubArguments = append(l.SubArguments, sarg)
	return &sarg
}

// NameExists accepts a proposed name for a
// sub-argument and checks if it already exists
// within the received Laywer.
func (l Lawyer) NameExists(n string) bool {
	// Format name
	n = strings.ToUpper(n)
	n = strings.TrimSpace(n)

	// Check if name exists
	for _, sa := range l.SubArguments {
		n2 := strings.ToUpper(sa.Name)
		if n == n2 {
			// Name found
			return true
		}
	}

	return false
}

// TakeCase implements TakeCustomCase with os.Args.
func (l Lawyer) TakeCase(mw bool) error {
	return l.TakeCustomCase(os.Args[1:], mw)
}

// TakeCustomCase accepts some arguments and will
// parse through them according to the sub-commands
// that the Lawyer has. The arguments passed to this
// function should not include the binary name.
func (l Lawyer) TakeCustomCase(arguments []string, mw bool) error {
	return nil
}
