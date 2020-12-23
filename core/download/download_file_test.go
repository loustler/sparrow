package download

import (
	"github.com/loustler/sparrow/core"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	err := DownloadFile("https://raw.githubusercontent.com/loustler/sparrow/main/LICENSE", core.Paths(core.HomePath(), "Downloads", "LICENSE"))

	if err != nil {
		t.Errorf("Cannot download LICENSE file for test. Because %s", err)
	}
}

func TestDownloadFileFromInvalidUrl(t *testing.T) {
	err := DownloadFile("https://127.0.0.1", "$HOME/failed")

	if err == nil {
		t.Fatalf("Unknown file downloads...")
	}
}