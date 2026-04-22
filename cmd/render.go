package cmd

import (
	"path"

	"github.com/rs/zerolog/log"

	"github.com/sachin-sankar/dyp/internal/core"
	"github.com/sachin-sankar/dyp/internal/parser"
	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:     "render",
	Short:   "Render a prompt from filename directly.",
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal().Msg("Too many arguments to render command. Expects prompt file name only.")
		}
		promptsDir, err := cmd.Flags().GetString("prompts")
		if err != nil {
			log.Fatal().Err(err).Msg("Error running render command.")
		}

		promptFile := args[0]
		prompt := parser.ParsePromptFile(path.Join(promptsDir, promptFile))
		core.RenderPrompt(prompt)
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
}
