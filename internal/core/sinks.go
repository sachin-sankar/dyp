package core

import (
	"net/url"
	"os/exec"
)

func WebSink(sink string, rendered string) {
	baseURLMap := map[string]string{"chatgpt": "https://chatgpt.com/?hints=search&q=", "perplexity": "https://www.perplexity.ai/search/?q=", "claude": "https://claude.ai/new?q="}
	finalURL := baseURLMap[sink] + url.QueryEscape(rendered)
	cmd := exec.Command("xdg-open", finalURL)
	cmd.Start()
}
