package fetch

import (
	"fmt"
	"time"
	"log"
	"runtime"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func OS() string {
	fetch_w := `
		$lastDir = pwd
		cd $HOME/.secman
		git pull
		cd $lastDir
	`
	
	fetch_ml := `
		cd ~/.secman
		git pull
		cd -
	`

	if runtime.GOOS == "windows" {
		return fetch_w
	} else {
		return fetch_ml
	}
}

func FetchSECDIR() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " ☄ Fetching..."
	s.Start()

	err, out, errout := shell.ShellOut(OS())

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	s.Stop()
	fmt.Print(out)
}
