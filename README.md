# Argue 2.x.x

A sassy Golang package for parsing command-line arguments.

**Note:** The legacy version of Argue is available on the 1.x.x branch of this repo.

## Installing

Run `go get github.com/rburmorrison/go-argue`.

## Usage

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
	agmt := argue.NewArgument("This is a test of the argument library.", "2.0.0")
	agmt.AddFlagFact("uint", "this is a uint", &tUInt)
	agmt.AddFlagFact("int", "this is an integer", &tInt)
	agmt.AddFlagFact("string", "this is a string", &tString)
	agmt.AddPositionalFact("pos", "this is a positional string", &tPos).SetRequired(true)
	agmt.AddPositionalFact("other", "this is another int", &tOther)

	// 3. Dispute the command-line arguments
	agmt.Dispute(os.Args[1:], true)

	// 4. Handle the results
	fmt.Println("tUInt:", tUInt)
	fmt.Println("tInt:", tInt)
	fmt.Println("tString:", tString)
	fmt.Println("tPos:", tPos)
	fmt.Println("tOther:", tOther)
}
```

## Bugs

If you come across any bugs while using argue, please submit an issue to this repo.

## Roadmap

- [ ] Add the ability to construct an argument from a struct