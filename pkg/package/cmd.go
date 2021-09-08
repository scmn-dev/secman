package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/scmn-dev/gh-api/git"
	"github.com/scmn-dev/secman/tools/packages"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/scmn-dev/gh-api/pkg/cmdutil"
	"github.com/scmn-dev/gh-api/utils"
	"github.com/scmn-dev/gh-api/core/ghrepo"
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
		&cobra.Command{
			Use:   "install <package>",
			Short: "Install a secman package",
			Args:  cmdutil.MinimumArgs(1, "must specify a package to install from"),
			RunE: func(cmd *cobra.Command, args []string) error {
				if args[0] == "." {
					wd, err := os.Getwd()
					if err != nil {
						return err
					}

					return m.InstallLocal(wd)
				}

				repo, err := ghrepo.FromFullName(args[0])
				if err != nil {
					return err
				}

				if err := checkValidPackage(cmd.Root(), m, repo.RepoName()); err != nil {
					return err
				}

				cls, err := f.Cluster()
				if err != nil {
					return err
				}

				protocol, _ := cls.Get(repo.RepoHost(), "git_protocol")
				return m.Install(ghrepo.FormatRemoteURL(repo, protocol), io.Out, io.ErrOut)
			},
		},
	)

	return &cmd
}

func checkValidPackage(rootCmd *cobra.Command, m packages.PackageManager, pkgName string) error {
	if !strings.HasPrefix(pkgName, "sm-") {
		return errors.New("package repository name must start with `sm-`")
	}

	commandName := strings.TrimPrefix(pkgName, "sm-")
	if c, _, err := rootCmd.Traverse([]string{commandName}); err != nil {
		return err
	} else if c != rootCmd {
		return fmt.Errorf("%q matches the name of a built-in command", commandName)
	}

	for _, ext := range m.List(false) {
		if ext.Name() == commandName {
			return fmt.Errorf("there is already an installed package that provides the %q command", commandName)
		}
	}

	return nil
}
