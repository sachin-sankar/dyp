package utils

import (
	"log"
	"os"
	"path"

	"github.com/sachin-sankar/dyp/internal/parser"
)

func ListPromptFiles() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to find user HOME directory")
	}
	promptsLocation := path.Join(home, ".prompts")
	promptFiles, readError := os.ReadDir(promptsLocation)
	if readError != nil {
		log.Fatal("Unable to read .prompts directory")
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
