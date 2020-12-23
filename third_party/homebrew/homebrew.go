package homebrew

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func InstallHomebrew(out, err io.Writer, in io.Reader) error {
	return os.RunCommand(out, err, in, "bash", "-c", `"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)
}