package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/scmn-dev/secman/cluster"
	"github.com/scmn-dev/secman/tools/packages"
	"github.com/scmn-dev/secman/tools/looksh"
	tcexe "github.com/Timothee-Cardoso/tc-exe"
)

type Manager struct {
	dataDir    func() string
	lookPath   func(string) (string, error)
	lookSh     func() (string, error)
	newCommand func(string, ...string) *exec.Cmd
}

func NewManager() *Manager {
	return &Manager{
		dataDir:    cluster.DataDir,
		lookPath:   tcexe.LookPath,
		lookSh:     looksh.Look,
		newCommand: exec.Command,
	}
}
