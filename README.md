# Argue 2.x.x

A sassy Golang package for parsing command-line arguments.

**Note:** The legacy version of Argue is available on the 1.x.x branch of this repo.

## Installing

Run `go get github.com/rburmorrison/go-argue`.

## Usage

### Basic

Creating an argument parser with argue takes four steps.

```go
package main

import (
	"fmt"
	"os"

	"github.com/rburmorrison/go-argue"
)

func main() {
	// 1. Create placeholder variables
	var tUInt uint
	var tInt int
	var tString = "Default text"
	var tPos string
	var tOther int

	// 2. Create your argument and facts
	agmt := argue.NewArgument("This is a test of the argument package.", "2.0.0") // or argue.NewEmptyArgument()
	agmt.AddFlagFact("uint", "this is a uint", &tUInt)
	agmt.AddFlagFact("int", "this is an integer", &tInt)
	agmt.AddFlagFact("bool", "this is a boolean", &tBool)
	agmt.AddFlagFact("string", "this is a string", &tString)
	agmt.AddPositionalFact("pos", "this is a positional string", &tPos)
	agmt.AddPositionalFact("other", "this is another int", &tOther).SetRequired(false)

	// 3. Dispute the command-line arguments
	agmt.Dispute(true)

	// 4. Handle the results
	fmt.Println("tUInt:", tUInt)
	fmt.Println("tInt:", tInt)
	fmt.Println("tBool:", tBool)
	fmt.Println("tString:", tString)
	fmt.Println("tPos:", tPos)
	fmt.Println("tOther:", tOther)
}
```

Usage information is automatically generated and can be viewed with `yourbinary --help`.

```
[user@localhost Desktop]$ yourbinary --help
yourbinary 2.0.0
This is a test of the argument library.

Usage: yourbinary [--int VALUE] [--string VALUE] [--uint VALUE] POS OTHER

Positional arguments:
  POS                  this is a positional string
  OTHER                this is another int

Flags:
  -b, --bool            this is a boolean
  -i, --int VALUE       this is an integer
  -s, --string VALUE    this is a string
  -u, --uint VALUE      this is a uint
  -h, --help            display this help and exit
  -v, --version         display version and exit
```

### Using a Struct

Argue now supports auto-generation of aruguments from a struct. This idea was inspired by [go-arg](https://github.com/alexflint/go-arg), but is treated as an optional add-on in Argue. Each field accepts three tags:

- **options**: accepts the values "required" and "positional" separated by commas
- **init**: accepts a letter to use as the initial for a fact
- **help**: the description of a fact to display in the argument's usage

All fields are assumed to be flags unless explicitly stated otherwise in the options.

**Example Usage**

```go
package main

import (
	"fmt"

	"github.com/rburmorrison/go-argue"
)

// 1. Define a struct
type example struct {
	Field1     int  `options:"required,positional" help:"this is field1"`
	BoolField2 bool `init:"a" help:"this is boolean2"`
}

func main() {
	// 2. Create an instance of the struct and generate
	//    an argument from it
	var e example
	agmt := argue.NewEmptyArgumentFromStruct(&e)

	// 3. Dispute the command-line arguments
	agmt.Dispute(true)

	// 4. Handle the results
	fmt.Println("e.Field1:", e.Field1)
	fmt.Println("e.Bool2:", e.BoolField2)
}
```

## Bugs

If you come across any bugs while using argue, please submit an issue to this repo.

## Roadmap

- [ ] Add the ability to construct an argument from a struct
- [ ] *Possibly* add support for sub-commands