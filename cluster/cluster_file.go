package cluster

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"runtime"

	yaml "gopkg.in/yaml.v3"
)

const (
	CONFIG_DIR      = "CONFIG_DIR"
	XDG_CONFIG_HOME = "XDG_CONFIG_HOME"
	XDG_STATE_HOME  = "XDG_STATE_HOME"
	XDG_DATA_HOME   = "XDG_DATA_HOME"
	APP_DATA        = "AppData" // (windows only)
	LOCAL_APP_DATA  = "LocalAppData" // (windows only)
)

// Cluster path precedence
// 1. CONFIG_DIR
// 2. XDG_CONFIG_HOME
// 3. AppData (windows only)
// 4. HOME
func ClusterDir() string {
	var path string
	if a := os.Getenv(CONFIG_DIR); a != "" {
		path = a
	} else if b := os.Getenv(XDG_CONFIG_HOME); b != "" {
		path = filepath.Join(b)
	} else if c := os.Getenv(APP_DATA); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, "Secman CLI")
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, ".sm-cluster")
	}

	// If the path does not exist try migrating cluster from default paths
	if !dirExists(path) && os.Getenv(CONFIG_DIR) == "" {
		_ = autoMigrateClusterDir(path)
	}

	return path
}

// State path precedence
// 1. XDG_CONFIG_HOME
// 2. LocalAppData (windows only)
// 3. HOME
func StateDir() string {
	var path string
	if a := os.Getenv(XDG_STATE_HOME); a != "" {
		path = filepath.Join(a)
	} else if b := os.Getenv(LOCAL_APP_DATA); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, "Secman CLI")
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".sm-cluster", "state")
	}

	// If the path does not exist try migrating state from default paths
	if !dirExists(path) {
		_ = autoMigrateStateDir(path)
	}

	return path
}

// Data path precedence
// 1. XDG_DATA_HOME
// 2. LocalAppData (windows only)
// 3. HOME
func DataDir() string {
	var path string
	if a := os.Getenv(XDG_DATA_HOME); a != "" {
		path = filepath.Join(a)
	} else if b := os.Getenv(LOCAL_APP_DATA); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, "Secman CLI")
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".sm-cluster", "share")
	}

	return path
}

var errSamePath = errors.New("same path")
var errNotExist = errors.New("not exist")

// If clusters exist then move them to newPath
func autoMigrateClusterDir(newPath string) error {
	path, err := os.UserHomeDir()
	if oldPath := filepath.Join(path, ".sm-cluster"); err == nil && dirExists(oldPath) {
		return migrateDir(oldPath, newPath)
	}

	return errNotExist
}

// If state file exist then move it to newPath
func autoMigrateStateDir(newPath string) error {
	path, err := os.UserHomeDir()
	if oldPath := filepath.Join(path, ".sm-cluster"); err == nil && dirExists(oldPath) {
		return migrateFile(oldPath, newPath, "state.yml")
	}

	return errNotExist
}

func migrateFile(oldPath, newPath, file string) error {
	if oldPath == newPath {
		return errSamePath
	}

	oldFile := filepath.Join(oldPath, file)
	newFile := filepath.Join(newPath, file)

	if !fileExists(oldFile) {
		return errNotExist
	}

	_ = os.MkdirAll(filepath.Dir(newFile), 0755)
	return os.Rename(oldFile, newFile)
}

func migrateDir(oldPath, newPath string) error {
	if oldPath == newPath {
		return errSamePath
	}

	if !dirExists(oldPath) {
		return errNotExist
	}

	_ = os.MkdirAll(filepath.Dir(newPath), 0755)
	return os.Rename(oldPath, newPath)
}

func dirExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

func fileExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}

func ClusterFile() string {
	return filepath.Join(ClusterDir(), "cluster.yml")
}

func HostsClusterFile() string {
	return filepath.Join(ClusterDir(), "hosts.yml")
}

func ParseDefaultCluster() (Cluster, error) {
	return parseCluster(ClusterFile())
}

func HomeDirPath(subdir string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	newPath := filepath.Join(homeDir, subdir)

	return newPath, nil
}

var ReadClusterFile = func(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, pathError(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

var WriteClusterFile = func(filename string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0771)
	if err != nil {
		return pathError(err)
	}

	cfgFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600) // cargo coded from setup
	if err != nil {
		return err
	}

	defer cfgFile.Close()

	_, err = cfgFile.Write(data)
	return err
}

var BackupClusterFile = func(filename string) error {
	return os.Rename(filename, filename+".bak")
}

func parseClusterFile(filename string) ([]byte, *yaml.Node, error) {
	data, err := ReadClusterFile(filename)
	if err != nil {
		return nil, nil, err
	}

	root, err := parseClusterData(data)
	if err != nil {
		return nil, nil, err
	}

	return data, root, err
}

func parseClusterData(data []byte) (*yaml.Node, error) {
	var root yaml.Node
	err := yaml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}

	if len(root.Content) == 0 {
		return &yaml.Node{
			Kind:    yaml.DocumentNode,
			Content: []*yaml.Node{{Kind: yaml.MappingNode}},
		}, nil
	}

	if root.Content[0].Kind != yaml.MappingNode {
		return &root, fmt.Errorf("expected a top level map")
	}

	return &root, nil
}

func isLegacy(root *yaml.Node) bool {
	for _, v := range root.Content[0].Content {
		if v.Value == "github.com" {
			return true
		}
	}

	return false
}

func migrateCluster(filename string) error {
	b, err := ReadClusterFile(filename)
	if err != nil {
		return err
	}

	var hosts map[string][]yaml.Node
	err = yaml.Unmarshal(b, &hosts)
	if err != nil {
		return fmt.Errorf("error decoding legacy format: %w", err)
	}

	cfg := NewBlankCluster()
	for hostname, entries := range hosts {
		if len(entries) < 1 {
			continue
		}
		mapContent := entries[0].Content
		for i := 0; i < len(mapContent)-1; i += 2 {
			if err := cfg.Set(hostname, mapContent[i].Value, mapContent[i+1].Value); err != nil {
				return err
			}
		}
	}

	err = BackupClusterFile(filename)
	if err != nil {
		return fmt.Errorf("failed to back up existing cluster: %w", err)
	}

	return cfg.Write()
}

func parseCluster(filename string) (Cluster, error) {
	_, root, err := parseClusterFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			root = NewBlankRoot()
		} else {
			return nil, err
		}
	}

	if isLegacy(root) {
		err = migrateCluster(filename)
		if err != nil {
			return nil, fmt.Errorf("error migrating legacy cluster: %w", err)
		}

		_, root, err = parseClusterFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to reparse migrated cluster: %w", err)
		}
	} else {
		if _, hostsRoot, err := parseClusterFile(HostsClusterFile()); err == nil {
			if len(hostsRoot.Content[0].Content) > 0 {
				newContent := []*yaml.Node{
					{Value: "hosts"},
					hostsRoot.Content[0],
				}
				restContent := root.Content[0].Content
				root.Content[0].Content = append(newContent, restContent...)
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return NewCluster(root), nil
}

func pathError(err error) error {
	var pathError *os.PathError
	if errors.As(err, &pathError) && errors.Is(pathError.Err, syscall.ENOTDIR) {
		if p := findRegularFile(pathError.Path); p != "" {
			return fmt.Errorf("remove or rename regular file `%s` (must be a directory)", p)
		}

	}
	return err
}

func findRegularFile(p string) string {
	for {
		if s, err := os.Stat(p); err == nil && s.Mode().IsRegular() {
			return p
		}
		newPath := filepath.Dir(p)
		if newPath == p || newPath == "/" || newPath == "." {
			break
		}
		p = newPath
	}
	return ""
}
