package homebrew

import (
	"io"

	"github.com/loustler/sparrow/core/os"
)

func InstallHomebrew(errOut io.Writer, in io.Reader) error {
	return os.RunCommandWithoutOutput(
		errOut,
		in,
		"bash",
		"-c",
		`"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`,
	)
}
