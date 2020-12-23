package cmd

import (
	config "github.com/loustler/sparrow/cmd/config/app/cmd"
	"github.com/loustler/sparrow/cmd/core"
	initCmd "github.com/loustler/sparrow/cmd/init/app/cmd"
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	cmds := core.NewRootCommand(os.Stdin, os.Stdout, os.Stderr)

	cmds.AddCommand(config.NewConfigCommand(), initCmd.NewInitCommand())

	return cmds
}
