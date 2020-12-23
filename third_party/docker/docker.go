package docker

import "github.com/loustler/sparrow/core/download"

func GetDockerDesktop() {
	download.DownloadFile("", "$HOME/downloads/Docker.dmg")
}
