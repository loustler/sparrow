package os

import (
	"bytes"
	"testing"
)

func TestIsRootUser(t *testing.T) {
	var output bytes.Buffer
	var error bytes.Buffer

	isRootUser, err := IsRootUser(&error, &output)

	t.Logf("User is RootUser? %t", isRootUser)

	if err != nil {
		t.Errorf("Cannot check user is root user. By %s", err)
	}
}
