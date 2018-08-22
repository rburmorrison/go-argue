# Argue 2.3.1

A sassy Golang package for parsing command-line arguments.

**Note:** The legacy version of Argue is available on the 1.x.x branch of this repo.

## Installing

Run `go get github.com/rburmorrison/go-argue`.

## Documentation

[https://godoc.org/github.com/rburmorrison/go-argue](https://godoc.org/github.com/rburmorrison/go-argue)

## Usage

### Basic

Creating an argument parser with Argue takes four steps.

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
	agmt := argue.NewArgument("This is a test of the argument package.", "x.x.x") // or argue.NewEmptyArgument()
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
yourbinary x.x.x
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
- **init**: accepts a letter to use as the initial for a fact or nothing for no initial
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

### Sub-Commands
When you have more than one argument, you might want to use a Lawyer to help you get them straight. Here is full-featured example of how to use a Lawyer:

```go
package main

import (
	"fmt"

	argue "github.com/rburmorrison/go-argue"
)

// 1. Create argument structs
type commOne struct {
	Field1     int  `options:"required,positional" help:"this is field1"`
	BoolField2 bool `init:"a" help:"this is boolean2"`
	Variable   bool `init:""`
}

type commTwo struct {
	Var1    int  `options:"required,positional" help:"this is var1"`
	Second2 bool `init:"a" help:"this is second2"`
	Change  bool `init:""`
}

func main() {
	// 2. Create the placeholder variables
	var one commOne
	var two commTwo
	var variable bool

	// 3. Create the Lawyer, add the arguments and flags,
	//    and add the function to run if a command is run
	//    by the user
	law := argue.NewLawyer("This is a test of the arugment package.", "x.x.x")
	law.AddFact("test", "this is just a test flag", &variable)
	law.AddArgumentFromStruct("one", "this is the first", &one).SetHandler(func(v interface{}) {
		////////////////////////////////////////////////////////
		// NOTE: v contains the struct that was used to auto- //
		// generate the argument. If the argument was created //
		// manually, nil will be provided instead.            //
		////////////////////////////////////////////////////////

		// 5. Handle the data
		data := v.(commOne)
		fmt.Println(data.BoolField2)
	})

	law.AddArgumentFromStruct("two", "this is the second", &two).SetHandler(func(v interface{}) {
		// 5. Handle the data
		data := v.(commTwo)
		fmt.Println(data.Var1)
	})

	// 4. Have the Lawyer take the case
	law.TakeCase(true)
}
```

The usage output for this code would look like this:

```
[user@localhost Desktop]$ yourbinary --help
yourbinary x.x.x
This is a test of the arugment package.

Usage: yourbinary [--test] COMMAND

Flags:
  -t, --test       this is just a test flag
  -h, --help       display this help and exit
  -v, --version    display version and exit

Commands:
  one    this is the first
  two    this is the second

Run 'yourbinary <command> --help' for details about a command.
```

Running `yourbinary <command> --help` will print the usage output of the argument that represents that command.
## Purpose

Why did I create Argue? After all, there are plenty of other [argument parsing packages](https://github.com/avelino/awesome-go#command-line) for Go out there. For me, the pacakges that I tried from this list had at least one of three problems. The first problem was that they were too verbose and cumbersome. When I am creating a command-line application, I want to spend as little time as possible on writing the code to parse arguments properly. The second problem was ugly usage output. The usage output, to me, is the most important part. I want my users to be able to understand how to use my tool without getting distracted by formatting misalignment. They should be able to see the output and know exactly where everything is. The third problem was the lack of sub-command support. Some packages were perfect, but I couldn't use them for all my projects because I couldn't scale them to use sub-commands.

I created Argue to solve these problems. Argue has clean output, support for auto-generating a parser from a struct (or manually, if you prefer), and simple-to-implement sub-commands. If you have any suggestions for Argue, be sure to submit an issue to this repo detailing what you'd like to see in the future.

## Bugs

If you come across any bugs while using Argue, please submit an issue to this repo.

## Roadmap

- [x] Add the ability to construct an argument from a struct
- [x] Add support for sub-commands