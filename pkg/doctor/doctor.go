package doctor

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/gosh"
	"github.com/abdfnx/looker"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/api"
	"github.com/scmn-dev/secman/constants"
	"github.com/spf13/viper"
)

var (
	_, err = looker.LookPath("scc")
	status = ""

	smVersionStatus = ""
	sccStatus = ""
	sccVersionStatus = ""
	secmanConfigStatus = ""
	outErr, out, errout = gosh.RunOutput("scc -v")
	configErr = viper.ReadConfig(bytes.NewBuffer(constants.SecmanConfig()))
)

func Doctor(buildVersion string) {
	if len(os.Args) > 1 {
		if (os.Args[1] == "doctor" || os.Args[1] == "check") {
			latestVersion := api.GetLatest("secman-cli", false)
			latestSCCVersion := api.GetLatest("scc", false)

			if err == nil {
				sccStatus = constants.Checkmark + "secman core cli is installed."
			} else {
				sccStatus = constants.X + "secman core cli is not installed."
			}

			if buildVersion != latestVersion {
				smVersionStatus = constants.X + "secman is not the latest version."
			} else {
				smVersionStatus = constants.Checkmark + "secman on the latest version."
			}

			if out != "" {
				out = out[:len(out)-1]
			}

			if outErr != nil {
				if strings.Contains(errout, "not") {
					sccVersionStatus = ""
				}
			} else {
				if latestSCCVersion == out {
					sccVersionStatus = constants.Checkmark + "secman core cli on the latest version."
				} else {
					sccVersionStatus = constants.X + "secman core cli is not on the latest version."
				}
			}

			viper.SetConfigType("json")

			if configErr == nil {
				secmanConfigStatus = constants.Checkmark + "secman config is found."
			} else {
				secmanConfigStatus = constants.X + "secman config is not found."
			}

			if smVersionStatus != "" {
				status += smVersionStatus
			}

			if sccStatus != "" {
				if status != "" {
					status += "\n" + sccStatus
				} else {
					status += sccStatus
				}
			}

			if sccVersionStatus != "" {
				if status != "" {
					status += "\n" + sccVersionStatus
				} else {
					status += sccVersionStatus
				}
			}

			if secmanConfigStatus != "" {
				if status != "" {
					status += "\n" + secmanConfigStatus
				} else {
					status += secmanConfigStatus
				}
			}

			fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Doctor") + "\n\n" + status))
		}
	}
}
