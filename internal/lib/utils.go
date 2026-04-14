package utils

import (
	"os"
	"path"

	"github.com/rs/zerolog/log"

	"github.com/sachin-sankar/dyp/internal/parser"
)

func ListPromptFiles() []string {
	promptsLocation := GetPromptsDirectory()
	promptFiles, readError := os.ReadDir(promptsLocation)
	if readError != nil {
		log.Fatal().Err(readError).Msgf("Unable to read .prompts directory.")
	}
	var resultPromptFiles []string
	for _, promptFile := range promptFiles {
		resultPromptFiles = append(resultPromptFiles, path.Join(promptsLocation, promptFile.Name()))
	}
	return resultPromptFiles
}

func ListPrompts() []parser.Prompt {
	var result []parser.Prompt

	for _, promptFile := range ListPromptFiles() {
		filePrompt := parser.ParsePromptFile(promptFile)
		result = append(result, filePrompt)
	}

	return result
}
func GetPromptsDirectory() string {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to read user's $HOME directory.")
	}
	promptsLocation := path.Join(home, ".prompts")
	return promptsLocation
}
