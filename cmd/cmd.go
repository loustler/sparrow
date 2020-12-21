package cmd

import (
	config "github.com/loustler/mac-suite/cmd/config/app/cmd"
	"github.com/loustler/mac-suite/cmd/core"
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	cmds := core.NewOrganicCommand(os.Stdin, os.Stdout, os.Stderr)

	cmds.AddCommand(config.NewConfigCommand())

	return cmds
}
