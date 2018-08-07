# Argue

A sassy Golang package for parsing command-line arguments.

**Note:** This project is a work in progress. See the roadmap below for details.

## Installing

Run `go get github.com/rburmorrison/go-argue/...`.

## Usage

Creating an argument parser with argue takes three steps.

```go
package main

import "github.com/rburmorrison/go-argue/pkg/argue"

func main() {
    // 1. Make a new argument and variables to store the
    //    parsed values in 
    var boolFlag1 bool
    var intPositional1 int
    agmt := argue.NewEmptyArgument()

    // 2. Supply facts (flags and positional arguments)
    //    to your argument
    agmt.AddFact("flag", "this is an example flag", &boolFlag1)
    agmt.AddFact("positional", "this is an example positional argument", &intPositional1).
        SetPositional(true).
        SetRequired(true)

    // 3. Propose your argument
    agmt.Propose()

    // Your variables are populated after Propose is run
    fmt.Println("boolFlag1:", boolFlag1)
    fmt.Println("intPositional1:", intPositional1)

    // Run "yourexe --help" to see the usage information
}
```

Notes:

- Currently, only int, float64, string, and boolean values are accepted by argue
- Currently, values passed to strings can not contain spaces, or they will be treated as separate items

For more information such as adding descriptions to your argument, check out the `godoc` documentation.

## Bugs

If you come across any bugs while using argue, please submit an issue to this repo.

## Roadmap

- [x] Add support for flag arguments
- [x] Add support for positional arguments
- [ ] Add support for using quotes when supplying a string value
- [ ] Make the propose function optionally return an error instead of always ending the program upon failing
- [ ] Add the ability to construct an argument from a struct
- [ ] Add support for more varibles types (`float32`, `uint`, `int64`)