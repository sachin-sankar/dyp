package cmd

import (
	"net/url"
	"os/exec"
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
		supportedSinks := []string{"chatgpt", "perplexity", "claude"}
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
			for _, sink := range supportedSinks {
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
		baseURLMap := map[string]string{"chatgpt": "https://chatgpt.com/?hints=search&q=", "perplexity": "https://www.perplexity.ai/search/?q=", "claude": "https://claude.ai/new?q="}

		for _, sink := range args {
			finalURL := baseURLMap[sink] + url.QueryEscape(rendered)
			cmd := exec.Command("xdg-open", finalURL)
			cmd.Start()
		}
	},
}

func init() {
	rootCmd.AddCommand(sinkCmd)

	sinkCmd.Flags().StringVarP(&promptFile, "prompt-file", "p", "", "Render a prompt file directly.")

}
