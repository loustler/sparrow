package oh_my_zsh

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func InstallOhMyZsh(out, err io.Writer, in io.Reader) error {
	return os.RunCommand(out, err, in,"sh" , "-c", `"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`)
}
