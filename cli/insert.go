package cli

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/v6/pkg/pipe/insert"
	"github.com/spf13/cobra"
)

func InsertCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert",
		Short: "Insert a password to your vault.",
		Long: "Insert a password to your vault.",
		Aliases: []string{"new", "add"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := tea.NewProgram(insert.Insert(&PwOpts)).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.Logins, "login", "l", false, "Insert a login password to your vault.")
	cmd.Flags().BoolVarP(&PwOpts.CreditCards, "credit-card", "c", false, "Insert a credit card to your vault.")
	cmd.Flags().BoolVarP(&PwOpts.Emails, "email", "e", false, "Insert a email to your vault.")
	cmd.Flags().BoolVarP(&PwOpts.Notes, "note", "n", false, "Insert a note to your vault.")
	cmd.Flags().BoolVarP(&PwOpts.Servers, "server", "s", false, "Insert a server to your vault.")
	cmd.Flags().BoolVarP(&PwOpts.AutoGenerate, "auto-generate", "g", false, "Auto generate a secure password for password field (Only works with Logins type).")

	return cmd
}
