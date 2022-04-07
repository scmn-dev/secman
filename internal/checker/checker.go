package checker

import (
	"fmt"
	"strings"

	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/api"
	"github.com/scmn-dev/secman/v6/cli/factory"
	"github.com/scmn-dev/secman/v6/constants"
)

func Check(buildVersion string) {
	cliFactory := factory.New()
	stderr := cliFactory.IOStreams.ErrOut

	s := lipgloss.NewStyle().PaddingLeft(2)
	primary := lipgloss.NewStyle().Foreground(lipgloss.Color(constants.PRIMARY_COLOR))
	yellow := lipgloss.NewStyle().Foreground(lipgloss.Color(constants.YELLOW_COLOR))
	gray := lipgloss.NewStyle().Foreground(lipgloss.Color(constants.SECONDARY_COLOR))

	latestVersion := api.GetLatest("secman-cli", true)
	isFromHomebrew := isUnderHomebrew()
	isFromUsrBinDir := isUnderUsr()
	isFromScoop := isUnderScoop()
	isFromAppData := isUnderAppData()

	var command = func() string {
		if isFromHomebrew {
			return "brew upgrade secman"
		} else if isFromUsrBinDir {
			return "curl -sL https://u.secman.dev | bash"
		} else if isFromScoop {
			return "scoop update secman"
		} else if isFromAppData {
			return "iwr -useb https://w.secman.dev | iex"
		}

		return ""
	}

	if buildVersion != latestVersion {
		fmt.Fprintf(stderr, s.Render("\n%s %s → %s"),
		yellow.Render("There's a new version of ") + primary.Render("secman") + yellow.Render(" is avalaible:"),
		primary.Render(buildVersion),
		primary.Render(latestVersion) + "\n")

		if command() != "" {
			fmt.Fprintf(stderr, yellow.Render("  To upgrade, run: %s"), gray.Render(command()) + "\n")
		}
	}
}

var secmanExe, _ = looker.LookPath("secman")

func isUnderHomebrew() bool {
	return strings.Contains(secmanExe, "brew")
}

func isUnderUsr() bool {
	return strings.Contains(secmanExe, "usr")
}

func isUnderAppData() bool {
	return strings.Contains(secmanExe, "AppData")
}

func isUnderScoop() bool {
	return strings.Contains(secmanExe, "scoop")
}
