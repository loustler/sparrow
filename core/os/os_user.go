package os

import (
	"io"
	"strconv"
)

func IsRootUser(errOut io.Writer, in io.Reader) (bool, error) {
	uidString, err := RunCommand(errOut, in, "id", "-u")

	if err != nil {
		return false, err
	}

	uid, err := strconv.Atoi(uidString[:len(uidString) - 1])

	return uid == 0, err
}