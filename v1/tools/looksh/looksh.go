// +build !windows

package looksh

import "os/exec"

func Look() (string, error) {
	return exec.LookPath("sh")
}
