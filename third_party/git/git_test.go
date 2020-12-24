package git

import (
	"bytes"
	"testing"
)

func TestSetGitUserName(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := SetGitUserName(&error, &in, "loustler")

	if err != nil {
		t.FailNow()
	}
}

func TestSetGitUserEmail(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := SetGitUserEmail(&error, &in, "dev.loustler@gmail.com")

	if err != nil {
		t.FailNow()
	}
}

func TestSetGitUseGpg(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := SetGitUseGpg(&error, &in, true)

	if err != nil {
		t.FailNow()
	}
}

func TestSetGitUseSignWhenCommit(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := SetGitUseSignWhenCommit(&error, &in)

	if err != nil {
		t.FailNow()
	}
}

func TestExecuteGitGlobalConfig(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := executeGitGlobalConfig(&error, &in, "committer.name", "loustler")

	if err != nil {
		t.FailNow()
	}
}

func TestExecuteGitGlobalConfigWithInvalidOption(t *testing.T) {
	var error bytes.Buffer
	var in bytes.Buffer

	err := executeGitGlobalConfig(&error, &in, "hello", "world")

	if err == nil {
		t.FailNow()
	}
}
