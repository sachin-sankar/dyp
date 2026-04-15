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
		if len(result) != 0 {
			for _, prompt := range result {
				log.Debug().Msgf("Compare prompt file titles: %s@%s | %s@%s", prompt.Title, prompt.FilePath, filePrompt.Title, filePrompt.FilePath)
				if prompt.Title == filePrompt.Title {
					log.Fatal().Str("Conflicting title", prompt.Title).Str("Conflicting Path A", prompt.FilePath).Str("Conflicting Path B", filePrompt.FilePath).Msgf("Duplicate prompt file with the same title found.")
				} else {
					result = append(result, filePrompt)
				}
			}
		} else {
			result = append(result, filePrompt)
		}
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
