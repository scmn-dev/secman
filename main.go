package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/mgutz/ansi"
	"github.com/scmn-dev/secman/v6/cli/factory"
	"github.com/scmn-dev/secman/v6/cli/secman"
	"github.com/scmn-dev/secman/v6/tools"
	"github.com/spf13/cobra"
)

var (
	version string
	buildDate string
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	runtime.LockOSThread()

	cliFactory := factory.New()
	hasDebug := os.Getenv("DEBUG") != ""
	stderr := cliFactory.IOStreams.ErrOut

	if !cliFactory.IOStreams.ColorEnabled() {
		surveyCore.DisableColor = true
	} else {
		surveyCore.TemplateFuncsWithColor["color"] = func(style string) string {
			switch style {
				case "white":
					if cliFactory.IOStreams.ColorSupport256() {
						return fmt.Sprintf("\x1b[%d;5;%dm", 38, 242)
					}

					return ansi.ColorCode("default")

				default:
					return ansi.ColorCode(style)
			}
		}
	}

	if len(os.Args) > 1 && os.Args[1] != "" {
		cobra.MousetrapHelpText = ""
	}

	RootCmd := secman.Execute(cliFactory, version, buildDate)

	if cmd, err := RootCmd.ExecuteC(); err != nil {
		if err == tools.SilentError {
			return exitError
		} else if tools.IsUserCancellation(err) {
			if errors.Is(err, terminal.InterruptErr) {
				fmt.Fprint(stderr, "\n")
			}

			return exitCancel
		}

		tools.PrintError(stderr, err, cmd, hasDebug)

		return exitError
	}

	if secman.HasFailed() {
		return exitError
	}

	return exitOK
}
