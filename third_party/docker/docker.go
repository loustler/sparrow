package docker

import "github.com/loustler/sparrow/core/download"

func GetDockerDesktop() {
	download.Download("", "$HOME/downloads/Docker.dmg")
}
