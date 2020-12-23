package main

import (
	"os"

	"github.com/loustler/sparrow/cmd"
)

func main() {
	cmd.NewCommand(os.Stdin, os.Stdout, os.Stderr).Execute()
}
