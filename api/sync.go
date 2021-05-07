package sync

import (
	"fmt"
	"log"
	"runtime"
	"time"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
	checker "github.com/secman-team/version-checker"
)

push_w :=
	`
		$lastDir = pwd
		cd ~/.secman

		if (Test-Path -path .git) {
			git add .
			git commit -m "new change"
			git push
		}

		cd $lastDir
	`

push_ml :=
	`
		cd ~/.secman
		git add .
		git commit -m "new secman password"
		git push
		cd -
	`
			

var (
	NewCmdStart = &cobra.Command{
		Use:   "start",
		Aliases: []string{"go", "."},
		Example: "secman sync start",
		Short: "Start Sync your passwords.",
		Run: func(cmd *cobra.Command, args []string) {
			w :=
				`
					$SM_GH_UN = git config user.name
					cd ~/.secman

					git init

					echo "# My secman passwords - $SM_GH_UN" >> ~/.secman\README.md

					secman repo create $SM_GH_UN/.secman -y --private

					git add .
					git commit -m "new .secman repo"
					git branch -M trunk
					git remote add origin https://github.com/$SM_GH_UN/.secman
					git push -u origin trunk

					cd $lastDir
				`
			ml :=
				`
					SM_GH_UN=$(git config user.name)
					cd ~/.secman
					git init

					echo "# My secman passwords - $SM_GH_UN" >> ~/.secman/README.md

					secman repo create $SM_GH_UN/.secman -y --private

					git add .
					git commit -m "new .secman repo"
					git branch -M trunk
					git remote add origin https://github.com/$SM_GH_UN/.secman
					git push -u origin trunk
					cd -
				`
			
			shell.SHCore(ml, w)
			checker.Checker()
		},
	}

	NewCmdClone = &cobra.Command{
		Use:   "clone",
		Aliases: []string{"cn", "/"},
		Short: CloneHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			w := 
				`
					$SM_GH_UN = git config user.name
					$clone=secman repo clone $SM_GH_UN/.secman ~/.secman

					if (Test-Path -path ~/.secman) {
						Remove-Item ~/.secman -Recurse -Force
						$clone
					} else {
						$clone
					}
					`
				ml :=
					`
						SM_GH_UN=$(git config user.name)
						clone="secman repo clone $SM_GH_UN/.secman ~/.secman"
	
						if [ -d ~/.secman ]; then
							rm -rf ~/.secman
							${clone}
						else
							${clone}
						fi
					`
				check_w :=
					`
						if (Test-Path -path ~/.secman) {
							Write-Host "cloned successfully"
						}
					`
	
				check_ml := `if [ -d ~/.secman ]; then echo "cloned successfully ✅"; fi`
	
				shell.SHCore(ml, w)
				shell.SHCore(check_ml, check_w)
				checker.Checker()
			},
		}
	
		NewCmdPush = &cobra.Command{
			Use:   "push",
			Aliases: []string{"ph"},
			Short: "Push The New Passwords in .secman .",
			Run: func(cmd *cobra.Command, args []string) {
			
			shell.SHCore(push_ml, push_w)
			checker.Checker()
		},
	}

	NewCmdPull = &cobra.Command{
		Use:   "pull",
		Aliases: []string{"pl"},
		Short: "Pull The New Passwords from :USERNAME/.secman .",
		Run: func(cmd *cobra.Command, args []string) {
			w := 
				`
					$lastDir = pwd
					cd ~/.secman
					git pull
					cd -
				`
			
			ml :=
				`
					cd ~/.secman
					git pull
					cd -
				`

			shell.SHCore(ml, w)
			checker.Checker()
		},
	}
)

func Sync(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync <command>",
		Short: "Sync Your Passwords.",
		Long:  `Sync Your Passwords, by create a private repo at :USERNAME/.secman`,
		Example: heredoc.Doc(`
			secman sync start
			secman sync clone
		`),
	}

	cmd.AddCommand(NewCmdStart)
	cmd.AddCommand(NewCmdClone)
	cmd.AddCommand(NewCmdPush)
	cmd.AddCommand(NewCmdPull)

	return cmd
}

func CloneHelp() string {
	const msg string = "Clone your .secman from your private repo at https://github.com/"
	repo := "/.secman ."
	cmd := "git config user.name"

	err, username, errout := shell.SHCore(cmd, cmd)

	uname := strings.TrimSuffix(username, "\n")

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	if uname != "" {
		return msg + uname + repo
	} else {
		return msg + ":USERNAME" + repo
	}
}

func PushSync() {
	const Syncing string = " 📮 Syncing..."

	if runtime.GOOS == "windows" {
		err, out, errout := shell.PWSLOut(
		`
			$directoyPath = "~/.secman/.git"

			if (Test-Path -path $directoyPath) {
				Write-Host "Reading from .secman folder..."
			}
		`)

		fmt.Print(out)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		} else if out != "" {
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Suffix = Syncing
			s.Start()

			shell.PWSLCmd(push_w)

			s.Stop()
		}
	} else {
		err, out, errout := shell.ShellOut(
		`
			if [ -d ~/.secman/.git ]; then
				echo "📖 Reading from .secman folder..."
			fi
		`)

		fmt.Print(out)

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		} else if out != "" {
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Suffix = Syncing
			s.Start()

			shell.ShellCmd(push_ml)

			s.Stop()
		}
	}
}
