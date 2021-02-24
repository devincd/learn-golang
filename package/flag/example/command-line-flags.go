// Go by Example: Command-Line Flags
//
// Command-line flags are a common way to specify options
// for command-line programs. For example, in wc -l the -l
// is a command-line flag.
//
// Go provides a flag package supporting basic command-line
//  flag parsing. We'll use the package to implement our
// example command-line program.
//
// Basic flag declarations are available for string, integer, and
// boolean options. Here we declare a string flag work with a
// default value "foo" and a short description. This
// flag.String function returns a string pointer (not a string
// value); we'll see how toi use this pointer below.
//
// This declares numb and fork flags, using a similar
// approach to the work flag.
//
// It's also possible to declare an option that uses an existing
// var declared elsewhere in the program. Note that we need
// to pass in a pointer ti the flag declaration function.
//
// Once all flags are declared, call flag.Parse() to execute
// the command-line parsing.
//
// Here we'll just dump out the parsed options and any
// trailing positional arguments. Note that we need to
// dereference the pointers with e.g. *wordPtr to get the
// actual option values.
//
// To experiment with the command-line flags program it's
// best to first compile it and then run the resulting binary
// directly
//
// Try out the built program by first giving it values for all
// flags.
//
// Note that if you omit flags they automatically take their
// default values.
//
// Trailing positional arguments can be provided after any
// flags.
//
// Note that the flag package requires all flags to appear
// before positional arguments (otherwise the flags will be
// interpreted as positional arguments).
//
// Use -h or --help flags to get automatically generated help
// text for the command-line program
//
// (Deprecated) If you provide a flag that wasn't specified to the flag
// package, the program will print an error message and
// show the help text again.
//
// If you provide a flag that wasn't specified to the flag
// package, the program will show the help text again.
//
// https://gobyexample.com/command-line-flags
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
