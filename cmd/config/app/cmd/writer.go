package cmd

import (
	gojson "encoding/json"
	"github.com/loustler/mac-suite/cmd/config/app/options"
)

func WriteToPath(options *options.ConfigOptions) {
	gojson.Marshal(ConfigOptionToConfig(options))
}

