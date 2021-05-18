package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/secman-team/gh-api/pkg/cmd/factory"
	"github.com/secman-team/shell"
	checker "github.com/secman-team/version-checker"
	"github.com/spf13/cobra"

	"github.com/secman-team/secman/pkg/api"
	"github.com/secman-team/secman/pkg/clean"
	"github.com/secman-team/secman/pkg/edit"
	"github.com/secman-team/secman/pkg/fetch"
	"github.com/secman-team/secman/pkg/open"
	"github.com/secman-team/secman/pkg/gen"
	"github.com/secman-team/secman/pkg/initialize"
	"github.com/secman-team/secman/pkg/insert"
	"github.com/secman-team/secman/pkg/pio"
	"github.com/secman-team/secman/pkg/show"
	"github.com/secman-team/secman/pkg/upgrade"
	"github.com/secman-team/secman/tools/auth"
	"github.com/secman-team/secman/tools/config"
	"github.com/secman-team/secman/tools/repo"
	commands "github.com/secman-team/secman/tools/constants"
)

var (
	copyPass bool
	version  string
	versionDate  string
	RootCmd  = &cobra.Command{
		Use:   "secman",
		Short: "Print the contents of the vault.",
		Long: `Print the contents of the vault. If you have
not yet initialized your vault, it is necessary to run
the init subcommand in order to create your secman
directory, and initialize your cryptographic keys.`,
		Run: func(cmd *cobra.Command, args []string) {
			if exists, _ := pio.PassFileDirExists(); exists {
				show.ListAll()
			} else {
				cmd.Help()
			}

			checker.Checker()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Aliases: []string{"ver"},
		Short: "Print the version of your secman binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("secman version " + version + " " + versionDate)
			fmt.Println("https://github.com/secman-team/secman/releases/tag/" + version)
			checker.Checker()
		},
	}

	verxCmd = &cobra.Command{
		Use: "verx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(version)
		},
	}

	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Clean your ~/.secman (delete it).",
		Run: func(cmd *cobra.Command, args []string) {
			clean.Clean()
			checker.Checker()
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize your secman vault.",
		Long:  "Initialize the .secman directory, and generate your secret keys.",
		Run: func(cmd *cobra.Command, args []string) {
			initialize.Init()
			checker.Checker()
		},
	}

	upgradeCmd = &cobra.Command{
		Use:     "upgrade",
		Aliases: []string{"upg"},
		Short:   "Upgrade your secman if there's a new release.",
		Run: func(cmd *cobra.Command, args []string) {
			upg.Upgrade()
		},
	}

	uninstallCmd = &cobra.Command{
		Use:     "uninstall",
		Aliases: []string{"un"},
		Short:   "Uninstall Your Secman.",
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS == "windows" {
				fmt.Println("run sm-upg uninstall")
			} else {
				shell.ShellCmd(commands.Uninstall())
			}
		},
	}

	insertCmd = &cobra.Command{
		Use:     "insert",
		Short:   "Insert a file or password in to your vault.",
		Example: "secman insert core/docker.com",
		Args:    cobra.RangeArgs(1, 2),
		Long: `Add a site to your password store. This site can optionally be a part
of a group by prepending a group name and slash to the site name.
Will prompt for confirmation when a site path is not unique.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 2 {
				path := args[0]
				filename := args[1]
				insert.File(path, filename)
			} else {
				pathName := args[0]
				insert.Password(pathName)
			}

			sync.PushSync()
			checker.Checker()
		},
	}

	showCmd = &cobra.Command{
		Use:     "show",
		Aliases: []string{"read"},
		Example: "secman show core/docker-password",
		Short:   "Print the password of a secman entry.",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			show.Site(path, copyPass)
			checker.Checker()
		},
	}

	generateCmd = &cobra.Command{
		Use:     "gen",
		Aliases: []string{"generate"},
		Short:   "Generate a secure password.",
		Example: "secman generate",
		Long: `Prints a randomly generated password. The length of this password defaults
to (24). If a password length is specified as greater than 2048 then generate
will fail.`,
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			pwlen := -1
			if len(args) != 0 {
				pwlenStr := args[0]
				pwlenint, err := strconv.Atoi(pwlenStr)
				if err != nil {
					pwlen = -1
				} else {
					pwlen = pwlenint
				}
			}

			pass := gen.Generate(pwlen)
			fmt.Println(pass)
			checker.Checker()
		},
	}

	findCmd = &cobra.Command{
		Use:     "find",
		Aliases: []string{"ls"},
		Example: "secman find code.com",
		Short:   "Find a site that contains the site-path.",
		Long: `Prints all sites that contain the site-path. Used to print just
one group or all sites that contain a certain word in the group or name.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			show.Find(path)
			checker.Checker()
		},
	}

	renameCmd = &cobra.Command{
		Use:     "rename",
		Short:   "Rename an entry in the password vault.",
		Example: "secman rename core/docker.com",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.Rename(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	editCmd = &cobra.Command{
		Use:     "edit",
		Aliases: []string{"update"},
		Short:   "Change the password of a site in the vault.",
		Example: "secman edit core/docker.com",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.Edit(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	removeCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Example: "secman remove core/docker.com",
		Short:   "Remove a site from the password vault by specifying the entire site-path.",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.RemovePassword(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	fetchCmd = &cobra.Command{
		Use:     "fetch",
		Example: "secman fetch",
		Short:   "Fetch if there is a new password/s in ~/.secman.",
		Run: func(cmd *cobra.Command, args []string) {
			fetch.FetchSECDIR()
			checker.Checker()
		},
	}

	syncCmd = sync.Sync()
	openCmd = open.Open(factory.New("x"), nil)

	// with github
	repoCmd = repox.Repo(factory.New("x"))
	authCmd = authx.Auth(factory.New("x"))
	configCmd = configx.Config(factory.New("x"))
)

func init() {
	RootCmd.AddCommand(authCmd)
	RootCmd.AddCommand(repoCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.AddCommand(cleanCmd)
	RootCmd.AddCommand(fetchCmd)
	RootCmd.AddCommand(findCmd)
	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(insertCmd)
	RootCmd.AddCommand(removeCmd)
	RootCmd.AddCommand(editCmd)
	RootCmd.AddCommand(renameCmd)
	RootCmd.AddCommand(showCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(upgradeCmd)
	RootCmd.AddCommand(verxCmd)
	RootCmd.AddCommand(uninstallCmd)
	RootCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(openCmd)
}

// main
func main() {
	RootCmd.Execute()
}
