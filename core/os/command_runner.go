package os

import (
	"io"
	"os/exec"
)

func RunCommand(errOut io.Writer, in io.Reader, cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)

	command.Stderr = errOut
	command.Stdin = in

	output, err := command.Output()

	return string(output), err
}

func RunCommandWithoutOutput(errOut io.Writer, in io.Reader, cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	command.Stderr = errOut
	command.Stdin = in

	return command.Run()
}
