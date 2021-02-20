package upg

import (
	"fmt"
	"os/exec"
)

func Upgrade() {
	cmd := exec.Command("verx", "--upg")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
