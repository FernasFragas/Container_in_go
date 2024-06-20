package smalliterm

import (
	prompt "github.com/c-bata/go-prompt"
	"small_projects_in_go/smalliterm/helpers"
)

func Main_Iterm() {
	promp := prompt.New(
		helpers.Executor,
		helpers.Completer,
		prompt.OptionPrefix(helpers.GetPrompt()),
		prompt.OptionLivePrefix(helpers.ChangeLivePrefix))

	promp.Run()
}
