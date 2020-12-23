package options

type ConfigOptions struct {
	// configuration file path
	ConfigFilePath string

	// In git-config, user.name. See https://git-scm.com/docs/git-config
	GitUserName string

	// In git-config user.email. See https://git-scm.com/docs/git-config
	GitUserEmail string

	// Use gpg when commit git with signing as default
	GitUserGpg bool

	// User name for your computer
	UserName string
}

func NewDefaultConfigOptions() ConfigOptions {
	return ConfigOptions{

	}
}
