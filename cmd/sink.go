package cmd

import (
	"path"

	"github.com/rs/zerolog/log"
	"github.com/sachin-sankar/dyp/internal/core"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/sachin-sankar/dyp/internal/parser"
	"github.com/sachin-sankar/dyp/internal/ui"
	"github.com/spf13/cobra"
)

var promptFile string

// sinkCmd represents the sink command
var sinkCmd = &cobra.Command{
	Use:     "sink",
	Aliases: []string{"s"},
	Short:   "Render a prompt template to a sink.",
	Long: `dyp allows you to render a commmand to a sink.
	Supported sinks:
	- chatgpt
	- perplexity
	- claude
You can pass in multiple sinks too :)
Web sinks will be opened using xdg-open. 
By default it will render a interactive list of available prompts to choose from. 
If you need to render a prompt file directly use --prompt-file`,
	Run: func(cmd *cobra.Command, args []string) {
		supportedWebSinks := []string{"chatgpt", "perplexity", "claude"}
		var (
			promptsDir string
			err        error
		)
		promptsDir, err = cmd.Flags().GetString("prompts")
		if err != nil {
			log.Fatal().Err(err).Msg("Error running list command.")
		}
		if len(args) < 1 {
			log.Fatal().Msg("Atleast one sink has to be defined.")
		}
		for _, arg := range args {
			found := false
			for _, sink := range supportedWebSinks {
				if arg == sink {
					found = true
				}
			}
			if !found {
				log.Fatal().Msgf("Unsupported sink %s", arg)
			}
		}
		var choosenPrompt parser.Prompt
		if promptFile == "" {
			choosenPrompt = ui.SelectPromptFileInteractively(promptsDir)
		} else {
			choosenPrompt = parser.ParsePromptFile(path.Join(utils.PromptDirectory(promptsDir), promptFile))
		}
		rendered := core.RenderPrompt(choosenPrompt)
		for _, arg := range args {
			for _, sink := range supportedWebSinks {
				if arg == sink {
					core.WebSink(arg, rendered)
					continue
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sinkCmd)

	sinkCmd.Flags().StringVarP(&promptFile, "prompt-file", "p", "", "Render a prompt file directly.")
}
