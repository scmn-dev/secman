package ios

import (
	"github.com/scmn-dev/gh-api/pkg/cmdutil"
	"github.com/scmn-dev/gh-api/context"
	"github.com/scmn-dev/gh-api/api"
	"github.com/scmn-dev/gh-api/core/ghrepo"

	"github.com/spf13/cobra"

	// commands
	aCmd "github.com/scmn-dev/gh-api/pkg/cmd/auth"
	cCmd "github.com/scmn-dev/gh-api/pkg/cmd/config"
	rCmd "github.com/scmn-dev/gh-api/pkg/cmd/repo"
)


func Auth(f *cmdutil.Factory) *cobra.Command {
	cmd := aCmd.NewCmdAuth(f)
	return cmd
}

func Config(f *cmdutil.Factory) *cobra.Command {
	cmd := cCmd.NewCmdCluster(f)
	return cmd
}

func Repo(f *cmdutil.Factory) *cobra.Command {
	repoResolvingCmdFactory := *f
	repoResolvingCmdFactory.BaseRepo = resolvedBaseRepo(f)
	
	cmd := rCmd.NewCmdRepo(&repoResolvingCmdFactory)

	return cmd
}

func resolvedBaseRepo(f *cmdutil.Factory) func() (ghrepo.Interface, error) {
	return func() (ghrepo.Interface, error) {
		httpClient, err := f.HttpClient()
		if err != nil {
			return nil, err
		}

		apiClient := api.NewClientFromHTTP(httpClient)

		remotes, err := f.Remotes()
		if err != nil {
			return nil, err
		}

		repoContext, err := context.ResolveRemotesToRepos(remotes, apiClient, "")
		if err != nil {
			return nil, err
		}

		baseRepo, err := repoContext.BaseRepo(f.IOStreams)
		if err != nil {
			return nil, err
		}

		return baseRepo, nil
	}
}
