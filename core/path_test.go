package core

import (
	"strings"
	"testing"
)

func TestHomePath(t *testing.T) {
	home := HomePath()

	if home == "" || home == "HOME" {
		t.Fatalf("Cannot find home directory path")
	}
}

func TestPaths(t *testing.T) {
	actual := Paths("root", "path", "to")
	expect := "root/path/to"

	if actual != expect {
		t.FailNow()
	}
}

func TestPathsWithEmptyStrings(t *testing.T) {
	actual := Paths("", "", "")
	expect := "//"

	if actual != expect {
		t.FailNow()
	}
}

func TestPathWithNoArgs(t *testing.T) {
	actual := Paths()
	expect := ""

	if actual != expect {
		t.FailNow()
	}
}

func TestPathFromHomePath(t *testing.T) {
	restPaths := []string{"hello", "world", "!"}

	actual := PathFromHomePath(restPaths...)
	expected := strings.Join(append([]string{HomePath()}, restPaths...), "/")

	if actual != expected {
		t.FailNow()
	}
}

func TestPathFromHomePathIfNoArgs(t *testing.T) {
	actual := PathFromHomePath()
	expect := HomePath()

	if actual != expect {
		t.FailNow()
	}
}

func TestPathFromHomePathIfEmptyStrings(t *testing.T) {
	actual := PathFromHomePath("", "", "")
	expect := HomePath() + "///"

	if actual != expect {
		t.FailNow()
	}
}
