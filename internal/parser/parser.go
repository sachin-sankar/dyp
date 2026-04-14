package parser

import (
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/adrg/frontmatter"
)

type Metadata struct {
	Title string `yaml:"title"`
}

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
		log.Fatal().Err(err).Msgf("Unable to read prompt file at path %s", promptFilePath)
	}
	var metadata Metadata
	bodyBytes, parseErr := frontmatter.MustParse(strings.NewReader(string(data)), &metadata)
	body := string(bodyBytes)
	if parseErr != nil {
		log.Fatal().Msgf("Unable to parse frontmatter for prompt file  at %s", promptFilePath)
	}
	var vars []Var
	re := regexp.MustCompile(`\{\{(.+?)\}\}`)
	matches := re.FindAllStringSubmatch(body, -1)
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
			log.Fatal().Msgf("Error while parsing prompt file %s. Field '%s' is undefined", promptFilePath, typeString)
		}
		matchedVar := Var{varParts[0], varType, match[1]}
		vars = append(vars, matchedVar)
	}
	prompt := Prompt{metadata.Title, vars, body}
	return prompt
}
