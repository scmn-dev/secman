package upd

import (
	"fmt"
	"os/exec"
)

func Update() {
	cmd := exec.Command("verx", "--upd")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
