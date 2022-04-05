package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abdfnx/gosh"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
	"github.com/scmn-dev/secman/v6/pkg/pipe/lister"
	"github.com/spf13/cobra"
)

func ListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all passwords.",
		Long: "List all passwords in your vault.",
		Aliases: []string{"."},
		RunE: func(cmd *cobra.Command, args []string) error {
			if PwOpts.ShowTreeView {
				st := shared.DefaultStyles()

				s := spinner.New(spinner.CharSets[11], 100 * time.Millisecond)
				s.Suffix = " 📡 Preparing & Getting data..."
				s.Start()

				err, out, errout := gosh.RunOutput("scc . -t")

				s.Stop()

				if err != nil {
					if strings.Contains(errout, "401") || strings.Contains(out, "401") {
						head := st.Error.Render("\n\nYour Authentication is Expired.")
						body := st.Subtle.Render(" Refresh your authentication via `secman auth refresh`.")

						fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Lister") + st.Wrap.Render(head + body)).String())

						os.Exit(0)
					}

					fmt.Println(errout)
				}

				fmt.Print(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Lister") + "\n\n" + st.Wrap.Render(out)).String())
			} else {
				if err := tea.NewProgram(lister.Lister(), tea.WithAltScreen()).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(1)
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&PwOpts.ShowTreeView, "tree-view", "t", false, "List passwords in tree view.")

	return cmd
}
