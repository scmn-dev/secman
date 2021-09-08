package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/scmn-dev/secman/pkg/root"

	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/scmn-dev/gh-api/api"
	"github.com/scmn-dev/gh-api/core/ghrepo"
	"github.com/scmn-dev/gh-api/pkg/cmd/factory"
	"github.com/scmn-dev/gh-api/pkg/cmdutil"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/scmn-dev/secman/tools/shared"
)

var (
	version string
	versionDate string
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
	cmdFactory := factory.New()
	hasDebug := os.Getenv("DEBUG") != ""
	stderr := cmdFactory.IOStreams.ErrOut

	if !cmdFactory.IOStreams.ColorEnabled() {
		surveyCore.DisableColor = true
	} else {
		surveyCore.TemplateFuncsWithColor["color"] = func(style string) string {
			switch style {
			case "white":
				if cmdFactory.IOStreams.ColorSupport256() {
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

	RootCmd := root.NewCmdRoot(cmdFactory, version, versionDate)

	cfg, _ := cmdFactory.Cluster()

	if host, err := cfg.DefaultHost(); err == nil {
		ghrepo.SetDefaultHost(host)
	}

	if cmd, err := RootCmd.ExecuteC(); err != nil {
		if err == cmdutil.SilentError {
			return exitError
		} else if cmdutil.IsUserCancellation(err) {
			if errors.Is(err, terminal.InterruptErr) {
				fmt.Fprint(stderr, "\n")
			}
			return exitCancel
		}

		shared.PrintError(stderr, err, cmd, hasDebug)

		var httpErr api.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == 401 {
			fmt.Fprintln(stderr, "hint: try authenticating with `secman auth login`")
		}

		return exitError
	}

	if root.HasFailed() {
		return exitError
	}

	return exitOK
}
