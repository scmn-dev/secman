package checker

import (
	"fmt"
	"os/exec"
)

// code
func Checker() {
	cmd := exec.Command("bash", "abdfnx/secman")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
