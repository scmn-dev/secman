package ui

import (
	"fmt"
	"path/filepath"

	"github.com/abdfnx/gosh"
	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/lipgloss"
	gapi "github.com/scmn-dev/get-latest/api"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/pkg/initx"
)

func Update() {
	if constants.SmErr == nil {
		UpdateMain()
	} else {
		initx.Init()
		UpdateMain()
	}
}

func UpdateMain() {
	smuiLatest := gapi.LatestWithArgs("scmn-dev/secman", "")
	url := "https://github.com/scmn-dev/secman/releases/download/" + smuiLatest + "/smui.zip"

	currentSMUIVersion, err := dfs.ReadFileContent(filepath.Join(constants.SMUIPath, "tag.txt"))
	message := ""

	if err != nil {
		initx.Init()
	}

	currentSMUIVersion = currentSMUIVersion[:len(currentSMUIVersion)-1]

	if smuiLatest != currentSMUIVersion {
		uCmd := fmt.Sprintf(`
			if [ -d %s/ui ]; then
				rm -rf %s/ui
				wget %s
				sudo chmod 755 smui.zip
				unzip -qq smui.zip
				mv ui %s/ui
				rm smui.zip
			fi
		`, constants.DotSecmanPath, constants.DotSecmanPath, url, constants.DotSecmanPath)

		wCmd := fmt.Sprintf(`
			if (Test-Path -path %s/ui) {
				Remove-Item %s/ui -Recurse -Force
				Invoke-WebRequest %s
				Expand-Archive smui.zip
				Move-Item -Path ui -Destination %s
				Remove-Item smui.zip -Recurse -Force
			}
		`, constants.DotSecmanPath, constants.DotSecmanPath, url, constants.DotSecmanPath)

		gosh.RunMulti(uCmd, wCmd)

		message = "SMUI upgraded successfully."
	} else {
		message = "SMUI on the latest version."
	}

	fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman UI") + "\n\n" + constants.Checkmark + message))
}
