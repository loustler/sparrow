package os

import (
	"bytes"
	"testing"
)

func TestRunListCommand(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	ls, err := RunCommand(&error, &output, "ls")

	t.Logf("ls results is output: %s", ls)

	if err != nil {
		t.Errorf("Failed to execute ls command. By %s", error.String())
	}
}
