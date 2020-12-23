package core

import (
	"github.com/spf13/cobra"
	"io"
)

func NewRootCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	return &cobra.Command {
		Use: "sparrow SUBCOMMAND",
		Short: "sparrow helps set-up dev env on macOS",
		Long: `
		sparrow helps set-up development environment on macOS

		Find more information at:
			https://github.com/loustler/sparrow`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}