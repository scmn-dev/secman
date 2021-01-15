package vm

import (
	"fmt"
	"os/exec"
)

func Main() {
	cmd := exec.Command("cgit", "--smd")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
