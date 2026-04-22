package ui

import (
	"charm.land/huh/v2"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/sachin-sankar/dyp/internal/parser"
)

func SelectPromptFileInteractively(promptsDir string) parser.Prompt {
	promptsDir = utils.PromptDirectory(promptsDir)
	prompts := utils.ListPrompts(promptsDir)
	var choosen int
	var options []huh.Option[int]
	for index, prompt := range prompts {
		options = append(options, huh.NewOption(prompt.Title, index))
	}
	huh.NewSelect[int]().Title("Choose a Prompt to Render").Options(options...).Value(&choosen).Run()

	choosenPrompt := prompts[choosen]
	return choosenPrompt
}
