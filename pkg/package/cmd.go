package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/scmn-dev/gh-api/pkg/cmdutil"
)

func PackageCmd(f *cmdutil.Factory) *cobra.Command {
	m := f.PackageManager
	io := f.IOStreams

	cmd := cobra.Command{
		Use:   "package",
		Short: "Manage secman package",
		Long: heredoc.Docf(`
			Secman Package are repositories that provide more additional apps.
		`, "`"),
		Aliases: []string{"pkg", "packages"},
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "list",
			Short: "List installed packages",
			Args:  cobra.NoArgs,
			Aliases: []string{"ls"},
			RunE: func(cmd *cobra.Command, args []string) error {
				cmds := m.List(true)
				if len(cmds) == 0 {
					return errors.New("no packages installed")
				}

				cs := io.ColorScheme()
				t := utils.NewTablePrinter(io)
				for _, c := range cmds {
					var repo string
					if u, err := git.ParseURL(c.URL()); err == nil {
						if r, err := ghrepo.FromURL(u); err == nil {
							repo = ghrepo.FullName(r)
						}
					}

					t.AddField(fmt.Sprintf("sm %s", c.Name()), nil, nil)
					t.AddField(repo, nil, nil)
					var updateAvailable string
					if c.UpdateAvailable() {
						updateAvailable = "Upgrade available"
					}

					t.AddField(updateAvailable, nil, cs.Green)
					t.EndRow()
				}

				return t.Render()
			},
		},
	)

	return &cmd
}
