package cmd

import (
	"io"

	"github.com/loustler/sparrow/cmd/core"
	initCmd "github.com/loustler/sparrow/cmd/init/app/cmd"
	"github.com/spf13/cobra"
)

func NewCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	command := core.NewRootCommand(in, out, err)

	command.AddCommand(initCmd.NewInitCommand())

	return command
}
