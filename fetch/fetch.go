package fetch

import (
	"fmt"
	"time"
	"runtime"
	"log"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

func FetchSECDIR() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " Fetching..."
	s.Start()

	err, out, errout := shell.ShellOut("")
	
	if runtime.GOOS == "windows" {
		err, out, errout = shell.PWSLOut("~/sm/secman-sync.ps1 --sm")
	} else {
		err, out, errout = shell.ShellOut("secman-sync pl")
	}
		
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Println(errout)
	}
		
	s.Stop()
	fmt.Println(out)
}
