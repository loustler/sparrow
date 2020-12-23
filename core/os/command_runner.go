package os

import (
	"io"
	"os/exec"
)

func RunCommand(out, errOut io.Writer, in io.Reader, cmd string, args ...string) error {
	commandPath, err := exec.LookPath(cmd)

	if err != nil {
		return err
	}

	command := &exec.Cmd {
		Path: commandPath,
		Args: args,
		Stdout: out,
		Stderr: errOut,
		Stdin: in,
	}

	return command.Run()
}
