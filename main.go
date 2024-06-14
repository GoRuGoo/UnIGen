package main

import (
	"os"

	"github.com/gorugoo/unigen/cli"
)

func main() {
	c := &cli.CLI{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	os.Exit(c.Execute(os.Args))
}
