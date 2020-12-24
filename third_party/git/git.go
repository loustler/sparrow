package git

import (
	"io"
	"strconv"

	"github.com/loustler/sparrow/core/os"
)

func SetGitUserName(errOut io.Writer, in io.Reader, username string) error {
	return executeGitGlobalConfig(errOut, in, "user.name", username)
}

func SetGitUserEmail(errOut io.Writer, in io.Reader, email string) error {
	return executeGitGlobalConfig(errOut, in, "user.email", email)
}

func SetGitUseGpg(errOut io.Writer, in io.Reader, useGpgWhenSigning bool) error {
	return executeGitGlobalConfig(errOut, in, "commit.gpgsign", strconv.FormatBool(useGpgWhenSigning))
}

func SetGitUseSignWhenCommit(errOut io.Writer, in io.Reader) error {
	return executeGitGlobalConfig(errOut, in, "alias.cis", "commit -S")
}

func executeGitGlobalConfig(errOut io.Writer, in io.Reader, key string, value string) error {
	return os.RunCommandWithoutOutput(errOut, in, "git", "config", "--global", key, value)
}
