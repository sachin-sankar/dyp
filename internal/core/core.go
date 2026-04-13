package core

import (
	"fmt"
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
		huh.NewText().
			Title(variable).
			Value(&currentAnswer).Run()
		answers = append(answers, answer{variable, currentAnswer})
	}
	rendered := prompt.Body
	for _, filledAnswer := range answers {
		rendered = strings.Replace(rendered, "{{"+filledAnswer.question+"}}", filledAnswer.answer, 1)
	}
	fmt.Print(rendered)
}
