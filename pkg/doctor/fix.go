package doctor

import (
	"fmt"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
)

func CommandStyle(cmd string) string {
	return lipgloss.NewStyle().Foreground(constants.GRAY_COLOR).SetString(cmd).String()
}

func Fix(buildVersion string) {
	var (
		bug1 string
		bug2 string
		bug3 string
		bug4 string
	)

	_, out, _ := gosh.RunOutput("scc -v")
	if out != "" {
		out = out[:len(out)-1]
	}

	if buildVersion != latestVersion {
		bug1 = "to upgrade run " + CommandStyle("`secman upgrade`") + " to download the latest version of secman."
	}

	if err != nil {
		bug2 = "to install secman core cli run " + CommandStyle("`npm i -g @secman/scc`")
	}

	if err == nil {
		if latestSCCVersion != out {
			bug3 = "to upgrade secman core cli to the latest version run " + CommandStyle("`npm update -g @secman/scc`")
		}
	}

	if configErr != nil {
		bug4 = "to initialize secman config run " + CommandStyle("`secman init`")
	}

	fixSteps := ""

	if bug1 != "" {
		fixSteps += bug1
	}

	if bug2 != "" {
		if fixSteps != "" {
			fixSteps += "\n" + bug2
		} else {
			fixSteps += bug2
		}
	}

	if bug3 != "" {
		if fixSteps != "" {
			fixSteps += "\n" + bug3
		} else {
			fixSteps += bug3
		}
	}

	if bug4 != "" {
		if fixSteps != "" {
			fixSteps += "\n" + bug4
		} else {
			fixSteps += bug4
		}
	}

	if fixSteps == "" {
		fixSteps = "there's nothing to fix, everything is good"
	}

	fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Doctor") + "\n\n" + fixSteps))
}
