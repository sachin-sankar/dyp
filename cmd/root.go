package cmd

import (
	"os"

	huh "charm.land/huh/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sachin-sankar/dyp/internal/core"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var promtpsDir string
var rootCmd = &cobra.Command{
	Use:   "dyp",
	Short: "Dynamically render prompts on the fly.",
	Run: func(cmd *cobra.Command, args []string) {
		prompts := utils.ListPrompts(promtpsDir)
		var choosen int
		var options []huh.Option[int]
		for index, prompt := range prompts {
			options = append(options, huh.NewOption(prompt.Title, index))
		}
		huh.NewSelect[int]().Title("Choose a Prompt to Render").Options(options...).Value(&choosen).Run()

		choosenPrompt := prompts[choosen]

		core.RenderPrompt(choosenPrompt)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose, err := cmd.PersistentFlags().GetBool("verbose")
		if err != nil {
			log.Fatal().Err(err).Msgf("Error running root command")
		}
		promtpsDir, err = cmd.PersistentFlags().GetString("prompts")
		if err != nil {
			log.Fatal().Err(err).Msgf("Error running root command")
		}
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		if promtpsDir == "default" {
			promtpsDir = utils.GetDefaultPromptsDirectory()
		}
		log.Debug().Msgf("Prompts directory location: %s", promtpsDir)
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("verbose", false, "Run dyp in verbose mode to observe debug logs.")
	rootCmd.PersistentFlags().String("prompts", "default", "Specify a custom directory path to look for prompts in. Defaults to $HOME/.prompts")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
