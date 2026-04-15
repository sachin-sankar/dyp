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
	log.Debug().Msgf("Found %d prompt files at %s", len(resultPromptFiles), promptsLocation)
	return resultPromptFiles
}

func ListPrompts(promptsLocation string) []parser.Prompt {
	var result []parser.Prompt
	var titles []string

	for _, promptFile := range ListPromptFiles(promptsLocation) {
		filePrompt := parser.ParsePromptFile(promptFile)
		for _, title := range titles {
			if title == filePrompt.Title {
				log.Fatal().Str("Conflicting title", title).Str("Conflicting Prompt File", promptFile).Msgf("Duplicate prompt file with the same title found.")
			}
		}
		titles = append(titles, filePrompt.Title)
		result = append(result, filePrompt)
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
