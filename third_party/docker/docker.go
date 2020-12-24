package docker

import (
	"github.com/loustler/sparrow/core"
	"github.com/loustler/sparrow/core/file"
)

func GetDockerDesktop() error {
	return file.Download(
		"https://desktop.docker.com/mac/stable/Docker.dmg",
		core.Paths(core.HomePath(), "downloads", "Docker.dmg"),
	)
}
