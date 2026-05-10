package core

import (
	"net/url"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func WebSink(sink string, rendered string) {
	baseURLMap := map[string]string{"chatgpt": "https://chatgpt.com/?hints=search&q=", "perplexity": "https://www.perplexity.ai/search/?q=", "claude": "https://claude.ai/new?q="}
	finalURL := baseURLMap[sink] + url.QueryEscape(rendered)
	cmd := exec.Command("xdg-open", finalURL)
	cmd.Start()
}

func TerminalSink(sink string, rendered string) {
	cmdMap := map[string]string{"opencode": "--prompt", "copilot": "-i", "gemini": "-i"}
	cmd := exec.Command(sink, cmdMap[sink], rendered)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug().Str("start string", cmd.String()).Msgf("Starting cmd sink")
	startErr := cmd.Run()
	if startErr != nil {
		log.Error().Err(startErr).Msg("Error occured trying to start a terminal sink.")
	}
}
