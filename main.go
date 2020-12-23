package mac_suite

import (
	"github.com/loustler/sparrow/cmd"
	"os"
)

func main() {
	cmd.NewCommand(os.Stdin, os.Stdout, os.Stderr).Execute()
}