package ui

import (
	"fmt"
	"log"
	"runtime"
	"net/http"

	"github.com/scmn-dev/secman/api"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/pkg/initx"
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
	initx.Init()

	handler := api.NewSFS(http.Dir(constants.SMUIPath), IndexHandler)

	smuiView := fmt.Sprintf(`
███████╗ ███╗   ███╗ ██╗   ██╗ ██╗
██╔════╝ ████╗ ████║ ██║   ██║ ██║
███████╗ ██╔████╔██║ ██║   ██║ ██║ PORT: %s        OS: %s
╚════██║ ██║╚██╔╝██║ ██║   ██║ ██║ HOST: http://localhost:%s
███████║ ██║ ╚═╝ ██║ ╚██████╔╝ ██║
╚══════╝ ╚═╝     ╚═╝  ╚═════╝  ╚═╝
`, constants.SMUI_PORT, runtime.GOOS, constants.SMUI_PORT)

	fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman UI") + "\n" + smuiView))

	log.Fatal(http.ListenAndServe(":" + constants.SMUI_PORT, handler))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, constants.SMUIIndexPath)
}
