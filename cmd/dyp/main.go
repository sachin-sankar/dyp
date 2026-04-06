// cmd/myapp/main.go
package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

type Prompt struct {
	title string
	vars  []string
}

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

func parsePromptFile(promptFilePath string) Prompt {
	data, err := os.ReadFile(promptFilePath)
	if err != nil {
		log.Fatal("Unable to read file " + promptFilePath)
	}
	source := string(data)
	parts := strings.Split(source, "---")
	title := parts[0]
	var vars []string
	re := regexp.MustCompile(`\{\{(.+?)\}\}`)
	matches := re.FindAllStringSubmatch(parts[1], -1)
	for _, match := range matches {
		vars = append(vars, match[1])
	}
	prompt := Prompt{title, vars}
	return prompt
}

func main() {
	prompt := (parsePromptFile(listPromptFiles()[0]))
	for _, i := range prompt.vars {
		fmt.Println(i)
	}
}
