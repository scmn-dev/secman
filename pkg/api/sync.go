package sync

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
	"github.com/spf13/cobra"
	commands "github.com/secman-team/secman/tools/constants"
	"github.com/secman-team/gh-api/pkg/cmd/factory"
	"github.com/secman-team/secman/tools/git_config"
)

var (
	NewCmdStart = &cobra.Command{
		Use:   "start",
		Aliases: []string{"go", "."},
		Example: "secman sync start",
		Short: "Start Sync your passwords.",
		Run: func(cmd *cobra.Command, args []string) {
			exCmd := commands.StartEX()

			shell.SHCore(commands.Start_ml(), commands.Start_w())
			shell.SHCore(exCmd, exCmd)
		},
	}

	NewCmdClone = &cobra.Command{
		Use:   "clone",
		Aliases: []string{"cn", "/"},
		Short: CloneHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			cloneCmd := commands.Clone()

			shell.SHCore(cloneCmd, cloneCmd)
			shell.SHCore(commands.Clone_check_ml(), commands.Clone_check_w())
		},
	}

	NewCmdPush = &cobra.Command{
		Use:   "push",
		Aliases: []string{"ph"},
		Short: "Push The New Passwords in ~/.secman .",
		Run: func(cmd *cobra.Command, args []string) {
			shell.SHCore(commands.Push_ml(), commands.Push_w())
		},
	}

	NewCmdPull = &cobra.Command{
		Use:   "pull",
		Aliases: []string{"pl"},
		Short: PullHelp(),
		Run: func(cmd *cobra.Command, args []string) {
			shell.SHCore(commands.Pull_ml(), commands.Pull_w())
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

	return cmd
}

const dotSecman string = "/.secman ."

func PullHelp() string {
	return git_config.GitConfig(factory.New("x"), "Pull The New Passwords from ", dotSecman)
}

func SyncHelp() string {
	return git_config.GitConfig(factory.New("x"), "Sync Your Passwords, by create a private repo at ", dotSecman)
}

func CloneHelp() string {
	return git_config.GitConfig(factory.New("x"), "Clone your .secman from your private repo at https://github.com/", dotSecman)
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
