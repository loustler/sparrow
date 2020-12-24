package file

import (
	"testing"

	"github.com/loustler/sparrow/core"
)

func TestDownload(t *testing.T) {
	err := Download("https://raw.githubusercontent.com/loustler/sparrow/main/LICENSE", core.Paths(core.HomePath(), "Downloads", "LICENSE"))

	if err != nil {
		t.Errorf("Cannot file LICENSE file for test. Because %s", err)
	}
}

func TestDownloadFromInvalidUrl(t *testing.T) {
	err := Download("https://127.0.0.1", "$HOME/failed")

	if err == nil {
		t.Fatalf("Unknown file downloads...")
	}
}
