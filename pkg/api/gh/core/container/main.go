package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	surveyCore "github.com/AlecAivazis/survey/v2/core"
	"github.com/AlecAivazis/survey/v2/terminal"
	tcexe "github.com/Timothee-Cardoso/tc-exe"
	colorable "github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
	"github.com/scmn-dev/secman/pkg/api/gh/api"
	"github.com/scmn-dev/secman/pkg/api/gh/core/build"
	"github.com/scmn-dev/secman/pkg/api/gh/core/config"
	"github.com/scmn-dev/secman/pkg/api/gh/core/expand"
	"github.com/scmn-dev/secman/pkg/api/gh/core/ghinstance"
	"github.com/scmn-dev/secman/pkg/api/gh/core/ghrepo"
	"github.com/scmn-dev/secman/pkg/api/gh/core/run"
	"github.com/scmn-dev/secman/pkg/api/gh/core/update"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/factory"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/root"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmdutil"
	"github.com/scmn-dev/secman/pkg/api/gh/utils"
	"github.com/spf13/cobra"
)

var updaterEnabled = ""

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	buildDate := build.Date
	buildVersion := build.Version

	updateMessageChan := make(chan *update.ReleaseInfo)
	go func() {
		rel, _ := checkForUpdate(buildVersion)
		updateMessageChan <- rel
	}()

	hasDebug := os.Getenv("DEBUG") != ""

	cmdFactory := factory.New()
	stderr := cmdFactory.IOStreams.ErrOut

	if spec := os.Getenv("GH_FORCE_TTY"); spec != "" {
		cmdFactory.IOStreams.ForceTerminal(spec)
	}

	if !cmdFactory.IOStreams.ColorEnabled() {
		surveyCore.DisableColor = true
	} else {
		// override survey's poor choice of color
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

	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	// terminal. With this, a user can clone a repo (or take other actions) directly from explorer.
	if len(os.Args) > 1 && os.Args[1] != "" {
		cobra.MousetrapHelpText = ""
	}

	rootCmd := root.NewCmdRoot(cmdFactory, buildVersion, buildDate)

	cfg, err := cmdFactory.Config()
	if err != nil {
		fmt.Fprintf(stderr, "failed to read configuration:  %s\n", err)
		return exitError
	}

	// TODO: remove after FromFullName has been revisited
	if host, err := cfg.DefaultHost(); err == nil {
		ghrepo.SetDefaultHost(host)
	}

	if cmd, err := rootCmd.ExecuteC(); err != nil {
		if err == cmdutil.SilentError {
			return exitError
		} else if cmdutil.IsUserCancellation(err) {
			if errors.Is(err, terminal.InterruptErr) {
				// ensure the next shell prompt will start on its own line
				fmt.Fprint(stderr, "\n")
			}
			return exitCancel
		}

		printError(stderr, err, cmd, hasDebug)
		

		if strings.Contains(err.Error(), "Incorrect function") {
			fmt.Fprintln(stderr, "You appear to be running in MinTTY without pseudo terminal support.")
			fmt.Fprintln(stderr, "To learn about workarounds for this error, run secman help mintty")
			return exitError
		}

		var httpErr api.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == 401 {
			fmt.Fprintln(stderr, "hint: try authenticating with `secman auth login`")
		}

		return exitError
	}

	cmd, _, err := rootCmd.Traverse(expandedArgs)

	if err != nil || cmd == rootCmd {
		originalArgs := expandedArgs
		isShell := false
		expandedArgs, isShell, err = expand.ExpandAlias(cfg, os.Args, nil)

		if err != nil {
			fmt.Fprintf(stderr, "failed to process aliases:  %s\n", err)
			return exitError
		}

		if hasDebug {
			fmt.Fprintf(stderr, "%v -> %v\n", originalArgs, expandedArgs)
		}

		if isShell {
			exe, err := tcexe.LookPath(expandedArgs[0])
			if err != nil {
				fmt.Fprintf(stderr, "failed to run external command: %s", err)
				return exitError
			}

			externalCmd := exec.Command(exe, expandedArgs[1:]...)
			externalCmd.Stderr = os.Stderr
			externalCmd.Stdout = os.Stdout
			externalCmd.Stdin = os.Stdin
			preparedCmd := run.PrepareCmd(externalCmd)

			err = preparedCmd.Run()
			if err != nil {
				var execError *exec.ExitError
				if errors.As(err, &execError) {
					return exitCode(execError.ExitCode())
				}
				fmt.Fprintf(stderr, "failed to run external command: %s", err)
				return exitError
			}

			return exitOK
		}
	}

	if root.HasFailed() {
		return exitError
	}

	return exitOK
}

func printError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		fmt.Fprintf(out, "error connecting to %s\n", dnsError.Name)
		if debug {
			fmt.Fprintln(out, dnsError)
		}
		fmt.Fprintln(out, "check your internet connection or githubstatus.com")
		return
	}

	fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, cmd.UsageString())
	}
}

func shouldCheckForUpdate() bool {
	if os.Getenv("GH_NO_UPDATE_NOTIFIER") != "" {
		return false
	}
	if os.Getenv("CODESPACES") != "" {
		return false
	}
	return updaterEnabled != "" && !isCI() && utils.IsTerminal(os.Stdout) && utils.IsTerminal(os.Stderr)
}

// based on https://github.com/watson/ci-info/blob/HEAD/index.js
func isCI() bool {
	return os.Getenv("CI") != "" || // GitHub Actions, Travis CI, CircleCI, Cirrus CI, GitLab CI, AppVeyor, CodeShip, dsari
		os.Getenv("BUILD_NUMBER") != "" || // Jenkins, TeamCity
		os.Getenv("RUN_ID") != "" // TaskCluster, dsari
}

func checkForUpdate(currentVersion string) (*update.ReleaseInfo, error) {
	if !shouldCheckForUpdate() {
		return nil, nil
	}

	client, err := basicClient(currentVersion)
	if err != nil {
		return nil, err
	}

	repo := updaterEnabled
	stateFilePath := filepath.Join(config.StateDir(), "state.yml")
	return update.CheckForUpdate(client, stateFilePath, repo, currentVersion)
}

// BasicClient returns an API client for github.com only that borrows from but
// does not depend on user configuration
func basicClient(currentVersion string) (*api.Client, error) {
	var opts []api.ClientOption
	if verbose := os.Getenv("DEBUG"); verbose != "" {
		opts = append(opts, apiVerboseLog())
	}
	opts = append(opts, api.AddHeader("User-Agent", fmt.Sprintf("GitHub CLI %s", currentVersion)))

	token, _ := config.AuthTokenFromEnv(ghinstance.Default())
	if token == "" {
		if c, err := config.ParseDefaultConfig(); err == nil {
			token, _ = c.Get(ghinstance.Default(), "oauth_token")
		}
	}
	if token != "" {
		opts = append(opts, api.AddHeader("Authorization", fmt.Sprintf("token %s", token)))
	}
	return api.NewClient(opts...), nil
}

func apiVerboseLog() api.ClientOption {
	logTraffic := strings.Contains(os.Getenv("DEBUG"), "api")
	colorize := utils.IsTerminal(os.Stderr)
	return api.VerboseLog(colorable.NewColorable(os.Stderr), logTraffic, colorize)
}
