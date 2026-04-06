// cmd/myapp/main.go
package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sachin-sankar/dyp/internal/parser"
)

func listPromptFiles() []string {
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

func main() {
	prompt := (parser.ParsePromptFile(listPromptFiles()[0]))
	for _, i := range prompt.Vars {
		fmt.Println(i)
	}
}
