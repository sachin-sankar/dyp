package core

import (
	"fmt"
	"regexp"
	"strings"

	"charm.land/huh/v2"
	"github.com/sachin-sankar/dyp/internal/parser"
)

type answer struct {
	question string
	answer   string
}

func RenderPrompt(prompt parser.Prompt) {
	var answers []answer
	var currentAnswer string

	for _, variable := range prompt.Vars {
		switch variable.VarType {

		case parser.TextField:
			huh.NewText().
				Title(variable.Question).
				Value(&currentAnswer).Run()
			answers = append(answers, answer{variable.Match, currentAnswer})
			currentAnswer = ""

		case parser.InputField:
			huh.NewInput().Title(variable.Question).Value(&currentAnswer).Run()
			answers = append(answers, answer{variable.Match, currentAnswer})
			currentAnswer = ""

		case parser.SelectField:
			re := regexp.MustCompile(`<(.*?)>`)
			optionsString := re.FindStringSubmatch(variable.Match)[1]
			optionStrings := strings.Split(optionsString, ",")
			var options []huh.Option[string]
			for _, optionString := range optionStrings {
				options = append(options, huh.NewOption(optionString, optionString))
			}
			huh.NewSelect[string]().Title(variable.Question).Options(options...).Value(&currentAnswer).Run()
			answers = append(answers, answer{variable.Match, currentAnswer})
			currentAnswer = ""

		case parser.MultiSelectField:
			re := regexp.MustCompile(`<(.*?)>`)
			optionsString := re.FindStringSubmatch(variable.Match)[1]
			optionStrings := strings.Split(optionsString, ",")
			var options []huh.Option[string]
			for _, optionString := range optionStrings {
				options = append(options, huh.NewOption(optionString, optionString))
			}
			var selected []string
			huh.NewMultiSelect[string]().Title(variable.Question).Options(options...).Value(&selected).Run()
			answers = append(answers, answer{variable.Match, strings.Join(selected, ",")})
			currentAnswer = ""
		}
	}

	rendered := prompt.Body
	for _, filledAnswer := range answers {
		rendered = strings.Replace(rendered, "{{"+filledAnswer.question+"}}", filledAnswer.answer, 1)
	}

	fmt.Print(rendered)
}
