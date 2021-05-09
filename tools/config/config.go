package configx

import (
	configCmd "github.com/secman-team/gh-api/pkg/cmd/config"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/secman-team/gh-api/api"
	"github.com/secman-team/gh-api/core/ghinstance"
	"net/http"
	"github.com/secman-team/gh-api/pkg/iostreams"
	"github.com/secman-team/gh-api/core/ghrepo"
)

func Config(f *cmdutil.Factory) *cobra.Command {
	cmd := configCmd.NewCmdConfig(f)

	return cmd
}

type ConfStruct struct {
	HttpClient func() (*http.Client, error)
	IO         *iostreams.IOStreams
	BaseRepo   func() (ghrepo.Interface, error)
}

func GitConfig(f *cmdutil.Factory) string {
	opts := ConfStruct{
		HttpClient: f.HttpClient,
	}

	httpClient, herr := opts.HttpClient()

	if herr != nil {
		return herr.Error()
	}

	apiClient := api.NewClientFromHTTP(httpClient)
	currentUser, cerr := api.CurrentLoginName(apiClient, ghinstance.Default())

	if cerr != nil {
		return cerr.Error()
	}
	
	if currentUser != "" {
		return currentUser
	} else {
		return ":USERNAME"
	}
}
