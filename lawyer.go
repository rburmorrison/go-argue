package argue

import (
	"os"
	"reflect"
	"strings"
)

// Lawyer represents an entity that can parse through
// sub-commands. A Layer is essentially a manager for
// multiple arguments.
type Lawyer struct {
	Description  string
	Version      string
	SubArguments []*SubArgument
	ShowDesc     bool
	ShowVersion  bool

	defaultArgument Argument
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

// AddFact adds a fact to the Lawyer. Facts may not
// be positional as that would conflict with the
// sub-arguments.
func (l *Lawyer) AddFact(name string, help string, v interface{}) *Fact {
	name = StandardizeFactName(name)
	if _, ok := l.defaultArgument.NameExists(name); ok {
		panic("argument: fact name already exits within this lawyer")
	}

	return l.defaultArgument.AddFlagFact(name, help, v)
}

// AddArgumentFromStruct offers a new argument to the
// Lawyer with the passed parameters: name, help, and
// the argument to add.
func (l *Lawyer) AddArgumentFromStruct(n string, h string, str interface{}) *SubArgument {
	arg := NewEmptyArgumentFromStruct(str)
	return l.AddArgument(n, h, arg)
}

// AddArgument offers a new argument to the Lawyer
// with the passed parameters: name, help, and the
// argument to add.
func (l *Lawyer) AddArgument(n string, h string, arg Argument) *SubArgument {
	if l.NameExists(n) {
		panic("argue: this name already exists as a sub-argument")
	}

	// Add command suffix
	arg.commandSuffix = strings.ToLower(n)

	// Make sub-argument
	var sarg SubArgument
	sarg.Name = n
	sarg.Help = h
	sarg.Argument = arg
	sarg.handler = nil

	l.SubArguments = append(l.SubArguments, &sarg)
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
	commandArgs := arguments

	// Extract all flags up to a command
	var flags []string
	for _, arg := range arguments {
		if flagReg.MatchString(arg) {
			flags = append(flags, arg)

			// Shave off this flag
			commandArgs = commandArgs[1:]
		} else {
			break
		}
	}

	// Check if --help or --version are present
	for _, f := range flags {
		if l.ShowVersion && (f == "-v" || f == "--version") {
			l.PrintVersion()
			os.Exit(0)
		}

		if f == "-h" || f == "--help" {
			l.PrintUsage()
			os.Exit(0)
		}
	}

	// Try to dispute the default flags
	err := l.defaultArgument.DisputeCustom(flags, mw)
	if err != nil {
		return err
	}

	// Check if command exists
	if len(commandArgs) == 0 {
		if mw {
			l.PrintError("no command specified")
		}

		return ErrNoCommand
	}

	// Check if command is specified with the Lawyer
	cmd := strings.ToUpper(commandArgs[0])
	var subArgument *SubArgument
	var commandSpecified bool
	for _, sa := range l.SubArguments {
		name := strings.ToUpper(sa.Name)
		if cmd == name {
			subArgument = sa
			commandSpecified = true
			break
		}
	}

	if !commandSpecified {
		if mw {
			l.PrintError("unknown command '" + commandArgs[0] + "'")
		}

		return ErrUnknownCommand
	}

	// Try to dispute appropriate command
	err = subArgument.Argument.DisputeCustom(commandArgs[1:], mw)
	if err != nil {
		return err
	}

	// Run the handler if it is specified
	if subArgument.handler != nil {
		if subArgument.Argument.baseStruct != nil {
			val := reflect.ValueOf(subArgument.Argument.baseStruct)
			subArgument.handler(&l, reflect.Indirect(val).Interface())
		} else {
			subArgument.handler(&l, nil)
		}
	}

	return nil
}
