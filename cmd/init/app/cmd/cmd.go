package cmd

import (
	prompt "github.com/loustler/sparrow/prompt/init"
	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init",
		Short: "Initialization",
		Long:  "Initialization",
		Run: func(cmd *cobra.Command, args []string) {
			prompt.NewInitPrompt()
		},
	}

	return command
}
