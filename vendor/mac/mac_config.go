package mac

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func SetHostName(out, errOut io.Writer, in io.Reader, hostname string) error {
	err := os.RunCommand(out, errOut, in, "hostname", hostname)

	if err != nil {
		return err
	}

	err = os.RunCommand(out, errOut, in, "scutil", "--set", "HostName", hostname)

	if err != nil {
		return err
	}

	return err
}

func ShowHiddenFiles(out, err io.Writer, in io.Reader) error {
	return os.RunCommand(out, err, in, "defaults", "write" , "com.apple.Finder", "AppleShowAllFiles", "YES")
}