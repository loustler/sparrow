package cmd

import "github.com/spf13/cobra"

func NewInitCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "init",
		Short: "Initialization",
		Long:  "Initialization",
	}

	return command
}
