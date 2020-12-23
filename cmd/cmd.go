package cmd

import (
	config "github.com/loustler/sparrow/cmd/config/app/cmd"
	"github.com/loustler/sparrow/cmd/core"
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	cmds := core.NewRootCommand(os.Stdin, os.Stdout, os.Stderr)

	cmds.AddCommand(config.NewConfigCommand())

	return cmds
}
