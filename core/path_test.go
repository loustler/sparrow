package core

import "testing"

func TestHomePath(t *testing.T) {
	home := HomePath()

	if home == "" || home == "HOME" {
		t.Fatalf("Cannot find home directory path")
	}
}
