package os

import (
	"bytes"
	"testing"
)

func TestRunListCommand(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	err := RunCommand(&output, &error, &output, "ls")

	t.Logf("ls results is %s", output.String())

	if err != nil {
		t.Errorf("Failed to execute ls command. By %s", error.String())
	}
}
