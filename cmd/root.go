package cmd

import (
	"fmt"
	"os"

	utils "github.com/sachin-sankar/dyp/internal/lib"
	"github.com/sachin-sankar/dyp/internal/parser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dyp",
	Short: "Dynamically render prompts on the fly.",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := (parser.ParsePromptFile(utils.ListPromptFiles()[0]))
		for _, i := range prompt.Vars {
			fmt.Println(i)
		}
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
