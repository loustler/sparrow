package core

import (
	"github.com/spf13/cobra"
	"io"
)

func NewOrganicCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	return &cobra.Command {
		Use: "organic SUBCOMMAND",
		Short: "organic set up dev env on macOS",
		Long: `
		organic setup development environment on macOS

		Find more information at:
			https://github.com/loustler/organic`,
		Run: runHelp,
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}