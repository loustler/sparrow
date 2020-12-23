package core

import (
	"os"
	"strings"
)

func Paths(path ...string) string {
	return strings.Join(path, "/")
}

func HomePath() string {
	return os.Getenv("HOME")
}
