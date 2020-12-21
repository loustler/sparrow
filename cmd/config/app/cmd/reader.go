package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConfigFile(path string) Config {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Errorf("Cannot find config file from %s.", path)
	}

	var config Config

	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Errorf("config file format is invalid")
	}

	return config
}
