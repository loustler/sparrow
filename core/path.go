package core

import (
	"os"
	"strings"
)

func Paths(path ...string) string {
	return strings.Join(path, "/")
}

func PathFromHomePath(paths ...string) string {
	return Paths(append([]string{HomePath()}, paths...)...)
}

func HomePath() string {
	return os.Getenv("HOME")
}
