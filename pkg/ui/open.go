package ui

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/api"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/pkg/initx"
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
