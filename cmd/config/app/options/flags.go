package options

import "github.com/spf13/pflag"

func AddKConfigFlags(fs *pflag.FlagSet, options *ConfigOptions) {
	addConfigPathFlag(fs, options)
	addGitConfigFlags(fs, options)
	addSystemConfigFlags(fs, options)
}

func addConfigPathFlag(fs *pflag.FlagSet, options *ConfigOptions) {
	fs.StringVar(&options.ConfigFilePath, ConfigFlag, "", "Modify configuration file path")
}

func addGitConfigFlags(fs *pflag.FlagSet, options *ConfigOptions) {
	fs.StringVar(&options.GitUserName, GitUserNameFlag, "", "Modify user.name for git")
	fs.StringVar(&options.GitUserEmail, GitUserEmailFlag, "", "Modify user.email for git")
	fs.BoolVar(&options.GitUserGpg, GitUserGpgFlag, false, "Modify git config use gpg as default")
}

func addSystemConfigFlags(fs *pflag.FlagSet, options *ConfigOptions) {
	fs.StringVar(&options.UserName, SystemUserNameFlag, "", "Modify system user name")
}