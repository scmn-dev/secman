package fetch

import (
	"fmt"
	"time"
	"log"

	"github.com/briandowns/spinner"
	"github.com/secman-team/shell"
)

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

func FetchSECDIR() {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Suffix = " Fetching..."
	s.Start()

	err, out, errout := shell.ShellOut(fetch_ml, fetch_w)
	
	shell.SHCore()
		
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
		
	s.Stop()
	fmt.Print(out)
}
