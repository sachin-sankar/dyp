package cmd

import (
	"fmt"
	"os"
	"strings"

	huh "charm.land/huh/v2"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/spf13/cobra"
)

type answer struct {
	question string
	answer   string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dyp",
	Short: "Dynamically render prompts on the fly.",
	Run: func(cmd *cobra.Command, args []string) {
		prompts := utils.ListPrompts()
		var choosen int
		var options []huh.Option[int]
		for index, prompt := range prompts {
			options = append(options, huh.NewOption(prompt.Title, index))
		}
		huh.NewSelect[int]().Title("Choose a Prompt to Render").Options(options...).Value(&choosen).Run()

		choosenPrompt := prompts[choosen]
		var answers []answer
		var currentAnswer string
		for _, variable := range choosenPrompt.Vars {
			huh.NewText().
				Title(variable).
				Value(&currentAnswer).Run()
			answers = append(answers, answer{variable, currentAnswer})
		}
		rendered := choosenPrompt.Body
		for _, filledAnswer := range answers {
			rendered = strings.Replace(rendered, "{{"+filledAnswer.question+"}}", filledAnswer.answer, 1)
		}
		fmt.Print(rendered)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
