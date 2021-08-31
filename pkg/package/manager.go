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
)

type Manager struct {
	dataDir    func() string
	lookPath   func(string) (string, error)
	lookSh     func() (string, error)
	newCommand func(string, ...string) *exec.Cmd
}
