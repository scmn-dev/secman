package cli

import (
	"fmt"
	"os"

	"github.com/abdfnx/gosh"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/v6/pkg/pipe/read"
	"github.com/spf13/cobra"
)

func ReadCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "Print the password of a secman entry.",
		Long: "Print the password of a secman entry.",
		Aliases: []string{"modify", "change"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				PwOpts.Password = args[0]
			}

			if PwOpts.ShowJsonView {
				passwordType := ""
				showHiddenFlag := ""

				if PwOpts.ShowHidden {
					showHiddenFlag = "-p"
				}

				if PwOpts.Logins {
					passwordType = "-l"
				} else if PwOpts.CreditCards {
					passwordType = "-c"
				} else if PwOpts.Emails {
					passwordType = "-e"
				} else if PwOpts.Notes {
					passwordType = "-n"
				} else if PwOpts.Servers {
					passwordType = "-s"
				}
				
				err, out, errout := gosh.RunOutput("scc read " + passwordType + " " + PwOpts.Password + " " + showHiddenFlag)

				if err != nil {
					fmt.Println(errout)
					os.Exit(2)
				} else {
					fmt.Print("\n" + out)
				}
			} else {
				if err := tea.NewProgram(read.Read(&PwOpts)).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(2)
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.Logins, "login", "l", false, "Read password from logins type.")
	cmd.Flags().BoolVarP(&PwOpts.CreditCards, "credit-card", "c", false, "Read password from credit cards type.")
	cmd.Flags().BoolVarP(&PwOpts.Emails, "email", "e", false, "Read password from emails type.")
	cmd.Flags().BoolVarP(&PwOpts.Notes, "note", "n", false, "Read password from notes type.")
	cmd.Flags().BoolVarP(&PwOpts.Servers, "server", "s", false, "Read password from servers type.")
	cmd.Flags().BoolVarP(&PwOpts.ShowHidden, "show-hidden", "p", false, "Show hidden values.")
	cmd.Flags().BoolVarP(&PwOpts.ShowJsonView, "json", "j", false, "Print password in JSON view.")

	return cmd
}
