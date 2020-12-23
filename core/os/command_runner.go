package os

import (
	"io"
	"os/exec"
)

func RunCommand(out, errOut io.Writer, in io.Reader, cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	command.Stdout = out
	command.Stderr = errOut
	command.Stdin = in

	return command.Run()
}
