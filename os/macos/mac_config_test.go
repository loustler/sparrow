package macos

import (
	"bytes"
	"testing"
)

func TestSetHostName(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	err := SetHostName(&error, &output, "loustler")

	t.Logf("output is %s", output.String())

	if err != nil {
		t.Errorf("Fail to set hostname as loustler. By %s", error.String())
	}
}

func TestShowHiddenFiles(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	err := ShowHiddenFiles(&error, &output)

	t.Logf("output is %s", output.String())

	if err != nil {
		t.Errorf("Fail to set show all hidden files. By %s", error.String())
	}
}
