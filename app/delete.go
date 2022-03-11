package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/pkg/pipe/delete"
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
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.Logins, "logins", "l", false, "Delete password from logins type.")
	cmd.Flags().BoolVarP(&PwOpts.CreditCards, "credit-cards", "c", false, "Delete password from credit cards type.")
	cmd.Flags().BoolVarP(&PwOpts.Emails, "emails", "e", false, "Delete password from emails type.")
	cmd.Flags().BoolVarP(&PwOpts.Notes, "notes", "n", false, "Delete password from notes type.")
	cmd.Flags().BoolVarP(&PwOpts.Servers, "servers", "s", false, "Delete password from servers type.")

	return cmd
}
