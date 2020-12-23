package homebrew

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func InstallHomebrew(errOut io.Writer, in io.Reader) error {
	_, err := os.RunCommand(errOut, in, "bash", "-c", `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)

	return err
}