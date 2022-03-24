package info

import (
	"fmt"

	"github.com/abdfnx/gosh"
	"github.com/abdfnx/looker"
	"github.com/scmn-dev/secman/api"
	"github.com/charmbracelet/glamour"
	"github.com/scmn-dev/secman/internal/config"
)

func Info(version string) {
	user := "`" + config.Config("config.name") + "`"

	if user == "``" {
		user = "`No User`"
	}

	var smcVersion string

	_, err := looker.LookPath("scc")

	if err != nil {
		smcVersion = "`Unkown`"
	} else {
		err, out, _ := gosh.RunOutput("scc -v")

		if err != nil {
			smcVersion = "`Unkown`"
		}

		smcVersion = out[:len(out)-1]
	}

	out1 := fmt.Sprintf(`# Secman CLI

* Version: %s
* Secman Core Version: %s
* Secman Core CLI Version: %s

Current User: %s

> Made with ❤️ by [secman](https://github.com/scmn-dev)`, version, api.GetLatest("secman-core", false), smcVersion, user)
	
	cli, err := glamour.Render(out1, "dark")

	if err != nil {
		fmt.Println("could not render info")
		return
	}

	cli = cli[:len(cli)-1]

	fmt.Print(cli)
}
