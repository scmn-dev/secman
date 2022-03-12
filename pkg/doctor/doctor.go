package doctor

import (
	"os"
	"fmt"
	"bytes"

	"github.com/abdfnx/gosh"
	"github.com/spf13/viper"
	"github.com/abdfnx/looker"
	"github.com/scmn-dev/secman/api"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
)

var (
	_, err = looker.LookPath("scc")

	smVersionStatus = ""
	sccStatus = ""
	sccVersionStatus = ""
	secmanConfigStatus = ""
	latestVersion = api.GetLatest("secman-cli", false)
	latestSCCVersion = api.GetLatest("scc", false)
	outErr, out, errout = gosh.RunOutput("scc -v")
	configErr = viper.ReadConfig(bytes.NewBuffer(constants.SecmanConfig()))
)

func Doctor(buildVersion string) {
	if len(os.Args) > 1 {
		if (os.Args[1] == "doctor" || os.Args[1] == "check") {
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
				fmt.Println(errout)
				os.Exit(0)
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

			fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Doctor") + "\n\n" + smVersionStatus + "\n" + sccStatus + "\n" + sccVersionStatus + "\n" + secmanConfigStatus))
		}
	}
}
