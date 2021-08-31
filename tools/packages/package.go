package packages

import (
	"io"
)

//go:generate moq -rm -out package_mock.go . Package
type Package interface {
	Name() string
	Path() string
	URL() string
	IsLocal() bool
	UpdateAvailable() bool
}

//go:generate moq -rm -out manager_mock.go . PackageManager
type PackageManager interface {
	List(includeMetadata bool) []Package
	Install(url string, stdout, stderr io.Writer) error
	InstallLocal(dir string) error
	Upgrade(name string, force bool, stdout, stderr io.Writer) error
	Remove(name string) error
	Dispatch(args []string, stdin io.Reader, stdout, stderr io.Writer) (bool, error)
	Create(name string) error
}
