package app

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/pkg/pipe/lister"
)

func ListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all passwords.",
		Long: "List all passwords in your vault.",
		Aliases: []string{"."},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := tea.NewProgram(lister.Lister(), tea.WithAltScreen()).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}
