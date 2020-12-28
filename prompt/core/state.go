package core

type LivePrefixState struct {
	prefix string
	enable bool
}

var state LivePrefixState

func init() {
	state = LivePrefixState{
		prefix: "Welcome sparrow",
		enable: false,
	}
}

func (state *LivePrefixState) UpdatePrefix(prefix string) {
	state.prefix = prefix
}

func (state *LivePrefixState) Disable() {
	state.enable = false
}

func (state *LivePrefixState) Enable() {
	state.enable = true
}
