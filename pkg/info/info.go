package info

import (
	"fmt"

	"github.com/abdfnx/gosh"
	"github.com/scmn-dev/secman/api"
	"github.com/charmbracelet/glamour"
	"github.com/scmn-dev/secman/internal/config"
)

func Info(version string) {
	err, smcVersion, _ := gosh.RunOutput("sc -v")
	user := "`" + config.Config("config.name") + "`"

	if user == "``" {
		user = "`No User`"
	}

	// remove the last line
	smcVersion = smcVersion[:len(smcVersion)-1]

	if err != nil {
		fmt.Println("could not get sc version")
		return
	}

	out1 := fmt.Sprintf(`# Secman CLI

* Version: %s
* Secman Core Version: %s
* Secman Core CLI Version: %s

Current User: %s

> Made with ❤️ by [secman](https://github.com/scmn-dev)`, version, api.GetLatestCore(), smcVersion, user)
	
	cli, err := glamour.Render(out1, "dark")

	if err != nil {
		fmt.Println("could not render info")
		return
	}

	cli = cli[:len(cli)-1]

	fmt.Print(cli)
}
