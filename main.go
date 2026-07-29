// Package main is the entry point for the dither command line tool.
package main

import (
	"fmt"
	"os"

	"github.com/theMagicalKarp/dithering/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// Diagnostics belong on stderr so that they stay out of a redirected
		// or piped stdout. Nothing useful is left to do if the write fails.
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
