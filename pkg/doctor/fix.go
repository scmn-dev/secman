package doctor

import (
	"fmt"

	"github.com/abdfnx/gosh"
	"github.com/abdfnx/looker"
	// "github.com/charmbracelet/lipgloss"
	// "github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/pkg/initx"
	"github.com/spf13/viper"
)

func Fix(buildVersion string) {
	_, npmErr := looker.LookPath("npm")

	if npmErr != nil {
		fmt.Println("npm is not installed, please install it first.")
	} else {
		if err != nil {
			gosh.Run("npm i -g @secman/sc@latest")
		}

		if err == nil {
			if latestSCVersion != out {
				gosh.Run("npm update -g @secman/sc@latest")
			}
		}

		viper.SetConfigType("json")
	
		if configErr != nil {
			initx.Init()
		}
	}
}
