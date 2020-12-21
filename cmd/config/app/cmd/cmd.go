package cmd

import (
	"github.com/loustler/mac-suite/cmd/config/app/options"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmds := &cobra.Command {
		Use: "config --flags",
		Short: "organic configuration",
		Long: `
			Modify organic config files using flags like "organic config --git-user=loustler --git-email=dev.loustler@gmail.com"`,
	}

	configOption := options.NewDefaultConfigOptions()

	options.AddKConfigFlags(cmds.Flags(), &configOption)
	cmds.MarkFlagRequired(options.SystemUserNameFlag)

	return cmds
}
