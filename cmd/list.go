package cmd

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/rs/zerolog/log"
	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available prompts.",
	Run: func(cmd *cobra.Command, args []string) {
		promtpsDir, err := cmd.Flags().GetString("prompts")
		if promptsDir == "$HOME/.prompts" {
			promptsDir = utils.GetDefaultPromptsDirectory()
		}
		log.Debug().Msgf("Prompts directory location: %s", promptsDir)
		if err != nil {
			log.Fatal().Err(err).Msgf("Error running list command.")
		}
		prompts := utils.ListPrompts(promtpsDir)
		table := table.New().Headers("Title", "File Path")
		table.BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")))
		for _, prompt := range prompts {
			table.Row(prompt.Title, prompt.FilePath)
		}
		lipgloss.Println(table)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
