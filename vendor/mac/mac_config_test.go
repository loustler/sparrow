package mac

import (
	"bytes"
	"testing"
)

func TestSetHostName(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	err := SetHostName(&output, &error, &output, "loustler")

	t.Logf("output is %s", output.String())

	if err != nil {
		t.Errorf("Failed to set host name.\n By %s", error.String())
	}
}

func TestShowHiddenFiles(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	err := ShowHiddenFiles(&output, &error, &output)

	t.Logf("output is %s", output.String())

	if err != nil {
		t.Errorf("Failed to set show hidden files.\n By %s", error.String())
	}
}