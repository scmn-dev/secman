package sync

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
	"github.com/abdfnx/shell"
	"github.com/spf13/cobra"
	commands "github.com/scmn-dev/secman-v1/tools/constants"
	git_config "github.com/scmn-dev/secman-v1/tools/config"
	"github.com/scmn-dev/secman-v1/tools/shared"
)

var username = git_config.GitConfig()

var (
	NewCmdStart = &cobra.Command{
		Use:   "start",
		Aliases: []string{"go", "."},
		Example: "secman sync start",
		Short: "Start Sync your passwords.",
		Run: func(cmd *cobra.Command, args []string) {
			if username != ":username" {
				exCmd := commands.StartEX()

				shell.SHCore(commands.Start_ml(), commands.Start_w())
				shell.SHCore(exCmd, exCmd)
			} else {
				shared.AuthMessage()
			}
		},
	}

	NewCmdClone = &cobra.Command{
		Use:   "clone",
		Aliases: []string{"cn", "/"},
		Short: CloneHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			if username != ":username" {
				cloneCmd := commands.Clone()

				shell.SHCore(cloneCmd, cloneCmd)
				shell.SHCore(commands.Clone_check_ml(), commands.Clone_check_w())
			} else {
				shared.AuthMessage()
			}
		},
	}

	NewCmdPush = &cobra.Command{
		Use:   "push",
		Aliases: []string{"ph"},
		Short: "Push The New Passwords in ~/.secman .",
		Run: func(cmd *cobra.Command, args []string) {
			if username != ":username" {
				shell.SHCore(commands.Push_ml(), commands.Push_w())
			} else {
				shared.AuthMessage()
			}
		},
	}

	NewCmdPull = &cobra.Command{
		Use:   "pull",
		Aliases: []string{"pl"},
		Short: PullHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			if username != ":username" {
				shell.SHCore(commands.Pull_ml(), commands.Pull_w())
			} else {
				shared.AuthMessage()
			}
		},
	}

	FetchClone = &cobra.Command{
		Use:   "pull",
		Short: PullHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			if username != ":username" {
				if runtime.GOOS == "windows" {
					shell.PWSLCmd(commands.Clone())
				} else {
					fmt.Println("This command isn't avaliable for this platform")
				}
			} else {
				shared.AuthMessage()
			}
		},
	}
)

func Sync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync <command>",
		Short: "Sync Your Passwords.",
		Long:  SyncHelp(),
		Example: heredoc.Doc(`
			secman sync start
			secman sync clone
		`),
	}

	cmd.AddCommand(NewCmdStart)
	cmd.AddCommand(NewCmdClone)
	cmd.AddCommand(NewCmdPush)
	cmd.AddCommand(NewCmdPull)
	cmd.AddCommand(FetchClone)

	return cmd
}

const dotSecman string = "/.secman ."

func PullHelp() string {
	return git_config.GitConfigWithMsg("Pull The New Passwords from ", dotSecman)
}

func SyncHelp() string {
	return git_config.GitConfigWithMsg("Sync Your Passwords, by create a private repo at ", dotSecman)
}

func CloneHelp() string {
	return git_config.GitConfigWithMsg("Clone your .secman from your private repo at https://github.com/", dotSecman)
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

			shell.PWSLCmd(commands.Push_w())

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

			shell.ShellCmd(commands.Push_ml())

			s.Stop()
		}
	}
}
