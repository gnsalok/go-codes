package clidesign

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitOK    = 0
	ExitUsage = 2
)

func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)
	name := fs.String("name", "", "name to greet")
	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}
	if *name == "" {
		fmt.Fprintln(stderr, "-name is required")
		return ExitUsage
	}
	fmt.Fprintf(stdout, "hello, %s\n", *name)
	return ExitOK
}
