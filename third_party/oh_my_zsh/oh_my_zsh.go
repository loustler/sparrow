package oh_my_zsh

import (
	"github.com/loustler/sparrow/core/os"
	"io"
)

func InstallOhMyZsh(err io.Writer, in io.Reader) error {
	_, error := os.RunCommand(err, in,"sh" , "-c", `"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`)

	return error
}
