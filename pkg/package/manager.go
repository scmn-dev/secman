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

	tcexe "github.com/Timothee-Cardoso/tc-exe"
	"github.com/gepis/sm-gh-api/core/config"
	"github.com/scmn-dev/secman/tools/looksh"
	"github.com/scmn-dev/secman/tools/packages"
)

var localPackageUpgradeError = errors.New("local packages can not be upgraded")

type Manager struct {
	dataDir    func() string
	lookPath   func(string) (string, error)
	lookSh     func() (string, error)
	newCommand func(string, ...string) *exec.Cmd
}

func NewManager() *Manager {
	return &Manager{
		dataDir:    config.DataDir,
		lookPath:   tcexe.LookPath,
		lookSh:     looksh.Look,
		newCommand: exec.Command,
	}
}

func (m *Manager) Dispatch(args []string, stdin io.Reader, stdout, stderr io.Writer) (bool, error) {
	if len(args) == 0 {
		return false, errors.New("too few arguments in list")
	}

	var exe string
	pkgName := args[0]
	forwardArgs := args[1:]

	_pkg, _ := m.list(false)
	for _, p := range _pkg {
		if p.Name() == pkgName {
			exe = p.Path()
			break
		}
	}

	if exe == "" {
		return false, nil
	}

	var externalCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// Dispatch all pkg calls through the `sh` interpreter to support executable files with a
		// shebang line on Windows.
		shExe, err := m.lookSh()
		if err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				return true, errors.New("the `sh.exe` interpreter is required. Please install Git for Windows and try again")
			}

			return true, err
		}

		forwardArgs = append([]string{"-c", `command "$@"`, "--", exe}, forwardArgs...)
		externalCmd = m.newCommand(shExe, forwardArgs...)
	} else {
		externalCmd = m.newCommand(exe, forwardArgs...)
	}

	externalCmd.Stdin = stdin
	externalCmd.Stdout = stdout
	externalCmd.Stderr = stderr
	return true, externalCmd.Run()
}

func (m *Manager) List(includeMetadata bool) []packages.Package {
	pack, _ := m.list(includeMetadata)
	return pack
}

func (m *Manager) list(includeMetadata bool) ([]packages.Package, error) {
	dir := m.installDir()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var results []packages.Package
	for _, f := range entries {
		if !strings.HasPrefix(f.Name(), "sm-") {
			continue
		}
		var remoteUrl string
		updateAvailable := false
		isLocal := false
		exePath := filepath.Join(dir, f.Name(), f.Name())

		if f.IsDir() {
			if includeMetadata {
				remoteUrl = m.getRemoteUrl(f.Name())
				updateAvailable = m.checkUpdateAvailable(f.Name())
			}
		} else {
			isLocal = true
			if !isSymlink(f.Mode()) {
				// if this is a regular file, its contents is the local directory of the pkg
				p, err := readPathFromFile(filepath.Join(dir, f.Name()))
				if err != nil {
					return nil, err
				}

				exePath = filepath.Join(p, f.Name())
			}
		}

		results = append(results, &Package{
			path:            exePath,
			url:             remoteUrl,
			isLocal:         isLocal,
			updateAvailable: updateAvailable,
		})
	}

	return results, nil
}

func (m *Manager) getRemoteUrl(pkg string) string {
	gitExe, err := m.lookPath("git")
	if err != nil {
		return ""
	}

	dir := m.installDir()
	gitDir := "--git-dir=" + filepath.Join(dir, pkg, ".git")
	cmd := m.newCommand(gitExe, gitDir, "config", "remote.origin.url")
	url, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(url))
}

func (m *Manager) InstallLocal(dir string) error {
	name := filepath.Base(dir)
	targetLink := filepath.Join(m.installDir(), name)
	if err := os.MkdirAll(filepath.Dir(targetLink), 0755); err != nil {
		return err
	}

	return makeSymlink(dir, targetLink)
}

func (m *Manager) Install(cloneURL string, stdout, stderr io.Writer) error {
	exe, err := m.lookPath("git")
	if err != nil {
		return err
	}

	name := strings.TrimSuffix(path.Base(cloneURL), ".git")
	targetDir := filepath.Join(m.installDir(), name)

	externalCmd := m.newCommand(exe, "clone", cloneURL, targetDir)
	externalCmd.Stdout = stdout
	externalCmd.Stderr = stderr
	return externalCmd.Run()
}

func (m *Manager) installDir() string {
	return filepath.Join(m.dataDir(), "packages")
}

func (m *Manager) checkUpdateAvailable(pkg string) bool {
	gitExe, err := m.lookPath("git")
	if err != nil {
		return false
	}

	dir := m.installDir()
	gitDir := "--git-dir=" + filepath.Join(dir, pkg, ".git")
	cmd := m.newCommand(gitExe, gitDir, "ls-remote", "origin", "HEAD")
	lsRemote, err := cmd.Output()

	if err != nil {
		return false
	}

	remoteSha := bytes.SplitN(lsRemote, []byte("\t"), 2)[0]
	cmd = m.newCommand(gitExe, gitDir, "rev-parse", "HEAD")
	localSha, err := cmd.Output()
	if err != nil {
		return false
	}

	localSha = bytes.TrimSpace(localSha)
	return !bytes.Equal(remoteSha, localSha)
}

func isSymlink(m os.FileMode) bool {
	return m&os.ModeSymlink != 0
}

// reads the product of makeSymlink on Windows
func readPathFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()
	b := make([]byte, 1024)
	n, err := f.Read(b)
	return strings.TrimSpace(string(b[:n])), err
}

func (m *Manager) Upgrade(name string, force bool, stdout, stderr io.Writer) error {
	exe, err := m.lookPath("git")
	if err != nil {
		return err
	}

	pack := m.List(false)
	if len(pack) == 0 {
		return errors.New("no packages installed")
	}

	someUpgraded := false
	for _, f := range pack {
		if name == "" {
			fmt.Fprintf(stdout, "[%s]: ", f.Name())
		} else if f.Name() != name {
			continue
		}

		if f.IsLocal() {
			if name == "" {
				fmt.Fprintf(stdout, "%s\n", localPackageUpgradeError)
			} else {
				err = localPackageUpgradeError
			}

			continue
		}

		var cmds []*exec.Cmd
		dir := filepath.Dir(f.Path())
		if force {
			fetchCmd := m.newCommand(exe, "-C", dir, "--git-dir="+filepath.Join(dir, ".git"), "fetch", "origin", "HEAD")
			resetCmd := m.newCommand(exe, "-C", dir, "--git-dir="+filepath.Join(dir, ".git"), "reset", "--hard", "origin/HEAD")
			cmds = []*exec.Cmd{fetchCmd, resetCmd}
		} else {
			pullCmd := m.newCommand(exe, "-C", dir, "--git-dir="+filepath.Join(dir, ".git"), "pull", "--ff-only")
			cmds = []*exec.Cmd{pullCmd}
		}

		if e := runCmds(cmds, stdout, stderr); e != nil {
			err = e
		}

		someUpgraded = true
	}

	if err == nil && !someUpgraded {
		err = fmt.Errorf("no package matched %q", name)
	}

	return err
}

func (m *Manager) Remove(name string) error {
	targetDir := filepath.Join(m.installDir(), "sm-" + name)
	if _, err := os.Lstat(targetDir); os.IsNotExist(err) {
		return fmt.Errorf("no package found: %q", targetDir)
	}

	return os.RemoveAll(targetDir)
}

func runCmds(cmds []*exec.Cmd, stdout, stderr io.Writer) error {
	for _, cmd := range cmds {
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
