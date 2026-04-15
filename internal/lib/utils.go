package utils

import (
	"os"
	"path"

	"github.com/rs/zerolog/log"

	"github.com/sachin-sankar/dyp/internal/parser"
)

func ListPromptFiles(promptsLocation string) []string {
	promptFiles, readError := os.ReadDir(promptsLocation)
	if readError != nil {
		log.Fatal().Err(readError).Msgf("Unable to read Prompts directory %s.", promptsLocation)
	}
	var resultPromptFiles []string
	for _, promptFile := range promptFiles {
		resultPromptFiles = append(resultPromptFiles, path.Join(promptsLocation, promptFile.Name()))
	}
	return resultPromptFiles
}

func ListPrompts(promptsLocation string) []parser.Prompt {
	var result []parser.Prompt

	for _, promptFile := range ListPromptFiles(promptsLocation) {
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
func GetDefaultPromptsDirectory() string {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to read user's $HOME directory.")
	}
	promptsLocation := path.Join(home, ".prompts")
	return promptsLocation
}
