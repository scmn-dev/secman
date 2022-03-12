package ui

import (
	"fmt"
	"log"
	"runtime"
	"net/http"

	"github.com/abdfnx/gosh"
	"github.com/scmn-dev/secman/api"
	"github.com/scmn-dev/secman/constants"
	gapi "github.com/scmn-dev/get-latest/api"
)

func Open() {
	port := "3750"
	smuiLatest := gapi.LatestWithArgs("david-tomson/smui", "")
	url := "https://github.com/david-tomson/smui/releases/download/" + smuiLatest + "/smui.zip"

	uCmd := fmt.Sprintf(`
		if ! [ -d %s/ui ]; then
			wget %s
			sudo chmod 755 smui.zip
			unzip smui.zip
			mv ui %s
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
`, port, runtime.GOOS, port)

	log.Fatal(http.ListenAndServe(":" + port, handler))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, constants.SMUIIndexPath)
}
