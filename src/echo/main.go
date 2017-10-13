// echo project main.go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdin

func echo(newline bool, sep string, args []string) error {
	fmt.Fprintf(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
