package cmd

import (
	gojson "encoding/json"
	"github.com/loustler/sparrow/cmd/config/app/options"
)

func WriteToPath(options *options.ConfigOptions) {
	gojson.Marshal(ConfigOptionToConfig(options))
}
