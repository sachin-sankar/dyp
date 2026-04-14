package parser

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Prompt struct {
	Title string
	Vars  []Var
	Body  string
}

type Var struct {
	Question string
	VarType  VarTypes
	Match    string
}

type VarTypes int

const (
	InputField VarTypes = iota
	TextField
	SelectField
	MultiSelectField
)

func ParsePromptFile(promptFilePath string) Prompt {
	data, err := os.ReadFile(promptFilePath)
	if err != nil {
		log.Fatal("Unable to read file " + promptFilePath)
	}
	source := string(data)
	parts := strings.Split(source, "---")
	if len(parts) != 2 {
		log.Fatal("Invalid Prompt file" + promptFilePath)
	}
	title := parts[0]
	var vars []Var
	re := regexp.MustCompile(`\{\{(.+?)\}\}`)
	matches := re.FindAllStringSubmatch(parts[1], -1)
	for _, match := range matches {
		varParts := strings.Split(match[1], "|")
		var varType VarTypes
		typeString := strings.Split(varParts[1], "<")[0]
		switch typeString {
		case "text":
			varType = TextField
		case "select":
			varType = SelectField
		case "multiselect":
			varType = MultiSelectField
		case "input":
			varType = InputField

		default:
			log.Fatal("Invalid field type while parsing prompt file " + promptFilePath + " " + match[1])
		}
		matchedVar := Var{varParts[0], varType, match[1]}
		vars = append(vars, matchedVar)
	}
	prompt := Prompt{title, vars, parts[1]}
	return prompt
}
