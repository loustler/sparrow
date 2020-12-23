package cmd

import (
	"github.com/loustler/sparrow/cmd/config/app/options"
	"time"
)

type ConfigMeta struct {
	configPath string
	updatedAt int64
}

type Config struct {
	updatedAt int64
}

func ConfigOptionToConfig(options *options.ConfigOptions) Config {
	return Config {

	}
}

func (meta *ConfigMeta) touch() {
	meta.updatedAt = nowEpoch()
}

func (config *Config) touch() {
	config.updatedAt = nowEpoch()
}

func nowEpoch() int64 {
	return time.Now().UnixNano() / 1000000
}