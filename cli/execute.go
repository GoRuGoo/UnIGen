package cli

import (
	"io"
)

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func (c *CLI) Execute(args []string) int {
	return 0
}
