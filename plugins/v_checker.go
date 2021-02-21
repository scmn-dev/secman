package checker

import (
	"fmt"
	"os/exec"
	"github.com/briandowns/spinner"
	"time"
)

func Checker() {
	fmt.Println("checking for updates...")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	time.Sleep(time.Second)
	s.Stop()
	cmd := exec.Command("verx", "--sm")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
