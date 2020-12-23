package cmd

import (
	"github.com/loustler/sparrow/cmd/config/app/options"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmds := &cobra.Command {
		Use: "config --flags",
		Short: "Modify configuration",
		Long: `
			Modify configuration files using flags like "sparrow config --git-user=loustler --git-email=dev.loustler@gmail.com"`,
	}

	configOption := options.NewDefaultConfigOptions()

	options.AddKConfigFlags(cmds.Flags(), &configOption)
	cmds.MarkFlagRequired(options.SystemUserNameFlag)

	return cmds
}
