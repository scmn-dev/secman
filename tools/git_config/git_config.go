package git_config

import (
	"os"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/secman-team/gh-api/api"
	"github.com/secman-team/gh-api/core/ghinstance"
	"net/http"
	"github.com/secman-team/gh-api/pkg/iostreams"
	"github.com/secman-team/gh-api/core/ghrepo"
	"github.com/secman-team/gh-api/pkg/cmd/factory"
)

type ConfStruct struct {
	HttpClient func() (*http.Client, error)
	IO         *iostreams.IOStreams
	BaseRepo   func() (ghrepo.Interface, error)
}

func GitConfig(f *cmdutil.Factory, msg1 string, msg2 string) string {
	opts := ConfStruct{
		HttpClient: f.HttpClient,
	}

	httpClient, herr := opts.HttpClient()

	if herr != nil {
		return herr.Error()
	}

	apiClient := api.NewClientFromHTTP(httpClient)
	currentUser, _ := api.CurrentLoginName(apiClient, ghinstance.Default())
	
	cmdFactory := factory.New("x")
	configRootCmd := NewCmdConfigRoot(cmdFactory)
	cfg, _ := cmdFactory.Config()

	expandedArgs := []string{}

	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	cmd, _, _ := configRootCmd.Traverse(expandedArgs)

	if cmdutil.IsAuthCheckEnabled(cmd) && !cmdutil.CheckAuth(cfg) {
		return msg1 + ":username" + msg2
	} else {
		return msg1 + currentUser + msg2
	}
}
