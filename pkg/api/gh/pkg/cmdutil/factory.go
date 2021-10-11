package cmdutil

import (
	"net/http"

	"github.com/scmn-dev/secman/pkg/api/gh/context"
	"github.com/scmn-dev/secman/pkg/api/gh/core/config"
	"github.com/scmn-dev/secman/pkg/api/gh/core/ghrepo"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/iostreams"
	"github.com/scmn-dev/secman/tools/packages"
)

type Browser interface {
	Browse(string) error
}

type Factory struct {
	IOStreams *iostreams.IOStreams
	Browser   Browser

	HttpClient func() (*http.Client, error)
	BaseRepo   func() (ghrepo.Interface, error)
	Remotes    func() (context.Remotes, error)
	Config     func() (config.Config, error)
	Branch     func() (string, error)

	PackageManager packages.PackageManager

	Executable string
}
