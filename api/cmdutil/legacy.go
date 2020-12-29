package cmdutil

import (
	"fmt"
	"os"

	"github.com/abdfnx/secman/v3/api/config"
)

// TODO: consider passing via Factory
// TODO: support per-hostname settings
func DetermineEditor(cf func() (config.Config, error)) (string, error) {
	editorCommand := os.Getenv("SM_EDITOR")
	if editorCommand == "" {
		cfg, err := cf()
		if err != nil {
			return "", fmt.Errorf("could not read config: %w", err)
		}
		editorCommand, _ = cfg.Get("", "editor")
	}

	return editorCommand, nil
}
