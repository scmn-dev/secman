package open

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"text/template"

	"github.com/MakeNowJust/heredoc"
	"github.com/secman-team/gh-api/api"
	"github.com/secman-team/gh-api/core/ghinstance"
	"github.com/secman-team/gh-api/core/ghrepo"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/secman-team/gh-api/pkg/iostreams"
	"github.com/secman-team/gh-api/pkg/markdown"
	"github.com/secman-team/gh-api/utils"
	"github.com/spf13/cobra"
)

type browser interface {
	Browse(string) error
}

type ViewOptions struct {
	HttpClient func() (*http.Client, error)
	IO         *iostreams.IOStreams
	BaseRepo   func() (ghrepo.Interface, error)
	Browser    browser

	RepoArg string
	Web     bool
	Branch  string
}

func Open() {}
