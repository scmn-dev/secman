package cmdutil

import (
	"net/http"

	"github.com/cli/cli/context"
	"github.com/abdfnx/secman/v3/api/config"
	"github.com/abdfnx/secman/v3/api/repox"
	"github.com/cli/cli/pkg/iostreams"
)

type Factory struct {
	IOStreams  *iostreams.IOStreams
	HttpClient func() (*http.Client, error)
	BaseRepo   func() (repox.Interface, error)
	Remotes    func() (context.Remotes, error)
	Config     func() (config.Config, error)
	Branch     func() (string, error)
}
