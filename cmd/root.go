package cmd

import (
	"os"

	huh "charm.land/huh/v2"
	"github.com/sachin-sankar/dyp/internal/core"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/spf13/cobra"
)

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

		core.RenderPrompt(choosenPrompt)
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
