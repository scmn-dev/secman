package fetch

import (
	"fmt"
	"time"
	"log"
	"runtime"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
	commands "github.com/secman-team/secman/tools/constants"
)

func OS() string {
	if runtime.GOOS == "windows" {
		return commands.Fetch_w()
	} else {
		return commands.Fetch_ml()
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
