package macos

import (
	"io"

	"github.com/loustler/sparrow/core/os"
)

func SetHostName(errOut io.Writer, in io.Reader, hostname string) error {
	err := os.RunCommandWithoutOutput(errOut, in, "hostname", hostname)

	if err != nil {
		return err
	}

	return os.RunCommandWithoutOutput(errOut, in, "scutil", "--set", "HostName", hostname)
}

func ShowHiddenFiles(errOut io.Writer, in io.Reader) error {
	return os.RunCommandWithoutOutput(errOut, in, "defaults", "write", "com.apple.Finder", "AppleShowAllFiles", "YES")
}
