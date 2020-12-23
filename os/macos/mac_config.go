package macos

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func SetHostName(errOut io.Writer, in io.Reader, hostname string) error {
	_, err := os.RunCommand(errOut, in, "hostname", hostname)

	if err != nil {
		return err
	}

	_, err = os.RunCommand(errOut, in, "scutil", "--set", "HostName", hostname)

	return err
}

func ShowHiddenFiles(errOut io.Writer, in io.Reader) error {
	_, err := os.RunCommand(errOut, in, "defaults", "write", "com.apple.Finder", "AppleShowAllFiles", "YES")

	return err
}