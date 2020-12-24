package docker

import "testing"

func TestGetDockerDesktop(t *testing.T) {
	err := GetDockerDesktop()

	if err != nil {
		t.FailNow()
	}
}
