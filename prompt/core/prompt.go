package core

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func DefaultPrompt(executor prompt.Executor, completer prompt.Completer, options ...prompt.Option) *prompt.Prompt {
	return prompt.New(
		executor,
		completer,
		PromptOptions(options...)...,
	)
}

func PromptOptions(options ...prompt.Option) []prompt.Option {
	defaultOptions := []prompt.Option{
		prompt.OptionInputBGColor(prompt.Black),
		prompt.OptionInputTextColor(prompt.White),
		prompt.OptionTitle("sparrow"),
		prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
			input := strings.ToLower(in)

			return breakline && (input == "quit" || input == "exit")
		}),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			return state.prefix, state.enable
		}),
	}

	return append(defaultOptions, options...)
}

func UpdateLivePrefix(prefix string, enable bool) {
	if enable {
		state.Enable()
		state.UpdatePrefix(prefix)
	} else {
		state.Disable()
	}
}
