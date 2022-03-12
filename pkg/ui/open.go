package ui

import (
	"fmt"
	"log"
	"runtime"
	"net/http"

	"github.com/abdfnx/gosh"
	"github.com/scmn-dev/secman/api"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/pkg/initx"
	gapi "github.com/scmn-dev/get-latest/api"
)

func Open() {
	if constants.SmErr == nil {
		OpenMain()
	} else {
		initx.Init()
		OpenMain()
	}
}

func OpenMain() {
	smuiLatest := gapi.LatestWithArgs("david-tomson/smui", "")
	url := "https://github.com/david-tomson/smui/releases/download/" + smuiLatest + "/smui.zip"

	uCmd := fmt.Sprintf(`
		if ! [ -d %s/ui ]; then
			wget %s
			sudo chmod 755 smui.zip
			unzip smui.zip
			mv ui %s/ui
			rm smui.zip
		fi
	`, constants.DotSecmanPath, url, constants.DotSecmanPath)

	wCmd := fmt.Sprintf(`
		if (-not (Test-Path -path %s/ui)) {
			Invoke-WebRequest %s
			Expand-Archive smui.zip
			Move-Item -Path ui -Destination %s
			Remove-Item smui.zip -Recurse -Force
		}
	`, constants.DotSecmanPath, url, constants.DotSecmanPath)

	gosh.RunMulti(uCmd, wCmd)

	handler := api.NewSFS(http.Dir(constants.SMUIPath), IndexHandler)

	fmt.Printf(`
███████╗ ███╗   ███╗ ██╗   ██╗ ██╗
██╔════╝ ████╗ ████║ ██║   ██║ ██║
███████╗ ██╔████╔██║ ██║   ██║ ██║ PORT %s         OS %s
╚════██║ ██║╚██╔╝██║ ██║   ██║ ██║ HOST http://localhost:%s
███████║ ██║ ╚═╝ ██║ ╚██████╔╝ ██║
╚══════╝ ╚═╝     ╚═╝  ╚═════╝  ╚═╝
`, constants.SMUI_PORT, runtime.GOOS, constants.SMUI_PORT)

	log.Fatal(http.ListenAndServe(":" + constants.SMUI_PORT, handler))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, constants.SMUIIndexPath)
}
