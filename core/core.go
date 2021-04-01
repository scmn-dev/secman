package main

import (
	"fmt"
	"runtime"
	"strconv"
	"log"

	"github.com/secman-team/shell"

	"github.com/secman-team/secman/api/sync"
	"github.com/secman-team/secman/edit"
	"github.com/secman-team/secman/fetch"
	"github.com/secman-team/secman/gen"
	"github.com/secman-team/secman/initialize"
	"github.com/secman-team/secman/insert"
	"github.com/secman-team/secman/pio"
	"github.com/secman-team/secman/plugins"
	"github.com/secman-team/secman/show"
	"github.com/secman-team/secman/upgrade"
	"github.com/secman-team/secman/clean"
	"github.com/spf13/cobra"
)

var (
	copyPass bool
	version  string
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
			fmt.Println(version + " " + runtime.GOOS + " " + runtime.GOARCH)
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
		Example: "secman show core/docker.com",
		Short:   "Print the password of a secman entry.",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			show.Site(path, copyPass)
			checker.Checker()
		},
	}

	start_syncCmd = &cobra.Command{
		Use:     "start-sync",
		Example: "secman start-sync",
		Short:   "Start Sync your passwords.",
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS == "windows" {
				shell.PWSLCmd("& ~/sm/secman-sync.ps1 sync")
			} else {
				shell.ShellCmd("secman-sync sync")
			}

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
)

func init() {
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
	RootCmd.AddCommand(start_syncCmd)
}

// main
func main() {
	mlChecker :=
		`
			end() {
				echo "after install dependencies, run secman again"
			}

			_cmd() {
				if ! [ -x "$(command -v $1)" ]; then
					echo "installing $1..."
					$2
					sudo chmod 755 /usr/local/bin/secman*
					sudo chmod 755 /usr/local/bin/cgit*
					sudo chmod 755 /usr/local/bin/verx*
					end
				fi
			}

			if ! [ -d /home/sm ]; then
				echo "sm folder was not found"
				echo "installing sm..."
				sudo git clone https://github.com/secman-team/sm /home/sm
				echo "installing ruby deps..."
				gem install colorize optparse
				end
			fi

			_cmd verx "sudo wget -P /usr/local/bin https://raw.githubusercontent.com/secman-team/verx/HEAD/verx"

			_cmd cgit "sudo wget -P /usr/local/bin https://raw.githubusercontent.com/secman-team/corgit/HEAD/cgit"

			_cmd secman-un "sudo wget -P /usr/local/bin https://raw.githubusercontent.com/secman-team/secman/HEAD/packages/secman-un"

			_cmd secman-sync "sudo wget -P /usr/local/bin https://raw.githubusercontent.com/secman-team/secman/HEAD/api/sync/secman-sync"
		`

	wCheck :=
		`
			$directoyPath="$HOME\sm";

			if(!(Test-Path -path $directoyPath)) {
				Write-Host "installing sm..."
				git clone https://github.com/secman-team/sm-win $directoyPath
				Write-Host "installing ruby deps..."
				gem install colorize optparse
				Invoke-WebRequest https://raw.githubusercontent.com/secman-team/tools/HEAD/sm.sh -outfile $directoyPath\sm.sh
				Write-Host "after install dependencies, run secman again"
			}
		`	

	if runtime.GOOS == "windows" {
		err, out, errout := shell.PWSLOut(wCheck)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Println(errout)
		} else if out != "" {
			fmt.Println("some of secman dependencies're not found, secman is going to fix it")
		} else {
			if out == "" {
				RootCmd.Execute()
			}
		}

		fmt.Println(out)

	} else {
		err, out, errout := shell.ShellOut(mlChecker)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Println(errout)
		} else if out != "" {
			fmt.Println("some of secman dependencies're not found, secman is going to fix it")
		} else {
			if out == "" {
				RootCmd.Execute()
			}
		}

		fmt.Println(out)
	}
}
