package checker

import (
	"fmt"
	"os/exec"
)

func Checker() {
	cmd := exec.Command("verx", "--sm")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
