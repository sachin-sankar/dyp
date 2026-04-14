package cmd

import (
	"path"

	"github.com/rs/zerolog/log"

	"github.com/sachin-sankar/dyp/internal/core"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/sachin-sankar/dyp/internal/parser"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render a prompt from filename directly.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal().Msgf("Too many arguments to render command. Expects prompt file name only.")
		}

		promptFile := args[0]
		prompt := parser.ParsePromptFile(path.Join(utils.GetPromptsDirectory(), promptFile))
		core.RenderPrompt(prompt)
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
}
