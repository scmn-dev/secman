package cli

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/v6/pkg/pipe/delete"
	"github.com/spf13/cobra"
)

func DeleteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Aliases: []string{"del", "rm", "remove"},
		Short: "Delete a password from your vault.",
		Long:  `Delete a password from your vault.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				PwOpts.Password = args[0]
			}

			if err := tea.NewProgram(delete.Delete(&PwOpts)).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(2)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.Logins, "login", "l", false, "Delete password from logins type.")
	cmd.Flags().BoolVarP(&PwOpts.CreditCards, "credit-card", "c", false, "Delete password from credit cards type.")
	cmd.Flags().BoolVarP(&PwOpts.Emails, "email", "e", false, "Delete password from emails type.")
	cmd.Flags().BoolVarP(&PwOpts.Notes, "note", "n", false, "Delete password from notes type.")
	cmd.Flags().BoolVarP(&PwOpts.Servers, "server", "s", false, "Delete password from servers type.")

	return cmd
}
