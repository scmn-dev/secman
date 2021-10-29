package looksh

import (
	"os"
	"path/filepath"

	tcexe "github.com/Timothee-Cardoso/tc-exe"
)

func Look() (string, error) {
	shPath, shErr := tcexe.LookPath("sh")
	if shErr == nil {
		return shPath, nil
	}

	gitPath, err := tcexe.LookPath("git")
	if err != nil {
		return "", shErr
	}

	gitDir := filepath.Dir(gitPath)

	// regular Git for Windows install
	shPath = filepath.Join(gitDir, "..", "bin", "sh.exe")
	if _, err := os.Stat(shPath); err == nil {
		return filepath.Clean(shPath), nil
	}

	// git as a scoop shim, if it was installed with scoop
	shPath = filepath.Join(gitDir, "..", "apps", "git", "current", "bin", "sh.exe")
	if _, err := os.Stat(shPath); err == nil {
		return filepath.Clean(shPath), nil
	}

	return "", shErr
}
