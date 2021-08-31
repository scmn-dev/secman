package pkg

import (
	"path/filepath"
	"strings"
)

type Package struct {
	path            string
	url             string
	isLocal         bool
	updateAvailable bool
}

func (e *Package) Name() string {
	return strings.TrimPrefix(filepath.Base(e.path), "sm-")
}

func (e *Package) Path() string {
	return e.path
}

func (e *Package) URL() string {
	return e.url
}

func (e *Package) IsLocal() bool {
	return e.isLocal
}

func (e *Package) UpdateAvailable() bool {
	return e.updateAvailable
}
