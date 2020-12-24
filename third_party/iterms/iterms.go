package iterms

import (
	"github.com/loustler/sparrow/core"
	"github.com/loustler/sparrow/core/file"
)

func SetIterms2Config() error {
	return file.Download(
		"https://raw.githubusercontent.com/loustler/sparrow/main/assets/iterms2/com.googlecode.iterm2.plist",
		core.PathFromHomePath("Library", "Preferences", "com.googlecode.iterm2.plist"),
	)
}
