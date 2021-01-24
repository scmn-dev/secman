package main

import (
	"fmt"
	"strconv"

	"github.com/secman-team/secman/v5/edit"
	"github.com/secman-team/secman/v5/gen"
	"github.com/secman-team/secman/v5/initialize"
	"github.com/secman-team/secman/v5/insert"
	"github.com/secman-team/secman/v5/pio"
	"github.com/secman-team/secman/v5/show"
	checker "github.com/secman-team/tools"
	"github.com/secman-team/secman/v5/update"
	"github.com/secman-team/secman/v5/api/vm"
	"github.com/secman-team/secman/v5/api/sync"
	"github.com/spf13/cobra"
)

var (
	copyPass bool
	version string
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
		Use:   "ver",
		Short: "Print the version of your secman binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
			checker.Checker()
		},
	}

	vmCmd = &cobra.Command{
		Use:   "vm",
		Short: "secman virtaul machine (with docker).",
		Run: func(cmd *cobra.Command, args []string) {
			vm.Main()
		},
	}

	verxCmd = &cobra.Command{
		Use:   "verx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
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

	updCmd = &cobra.Command{
		Use:   "upd",
		Short: "Update your secman if there's a new release.",
		Run: func(cmd *cobra.Command, args []string) {
			upd.Update()
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

			checker.Checker()
			sync.PushSync()
		},
	}

	showCmd = &cobra.Command{
		Use:     "show",
		Example: "secman show core/docker.com",
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
			pass := generate.Generate(pwlen)
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
			checker.Checker()
			sync.PushSync()
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
			checker.Checker()
			sync.PushSync()
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
			checker.Checker()
			sync.PushSync()
		},
	}
)

func init() {
	showCmd.PersistentFlags().BoolVarP(&copyPass, "copy", "c", false, "Copy your password to the clipboard")
	RootCmd.AddCommand(findCmd)
	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(insertCmd)
	RootCmd.AddCommand(removeCmd)
	RootCmd.AddCommand(editCmd)
	RootCmd.AddCommand(renameCmd)
	RootCmd.AddCommand(showCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(updCmd)
	RootCmd.AddCommand(vmCmd)
	RootCmd.AddCommand(verxCmd)
}

// main
func main() {
	RootCmd.Execute()
}
