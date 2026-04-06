package parser

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Prompt struct {
	Title string
	Vars  []string
}

func ParsePromptFile(promptFilePath string) Prompt {
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
