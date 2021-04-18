package clone

import (
	"fmt"
	"log"
	"strings"
	"runtime"
	"github.com/secman-team/shell"
)

func Core() {
	if runtime.GOOS == "windows"{
		shell.PWSLCmd("& $HOME/sm/secman-sync.ps1 cn")
		shell.PWSLCmd(
			`
				if (Test-Path -path $HOME/.secman) {
					Write-Host "cloned successfully"
				}
			`,
		)
	} else {
		shell.ShellCmd("secman-sync cn")
		shell.ShellCmd(`if [ -d ~/.secman ]; then echo "cloned successfully ✅"; fi`)
	}
}

func Help() string {
	const msg string = "Clone your .secman from your private repo at https://github.com/"
	repo := "/.secman ."

	if runtime.GOOS == "windows" {
		err, username, errout := shell.PWSLOut("git config user.name")

		uname := strings.TrimSuffix(username, "\n")
	
		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		}

		if uname != "" {
			return msg + uname + repo
		} else {
			return msg + ":USERNAME" + repo
		}

	} else {
		err, username, errout := shell.ShellOut("git config user.name")

		uname := strings.TrimSuffix(username, "\n")

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		}

		if uname != "" {
			return msg + uname + repo
		} else {
			return msg + ":USERNAME" + repo
		}
	}
}
