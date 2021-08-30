package shared

import (
	"os"
	"fmt"
	"io"
	"errors"
	"net"
	"strings"

	"github.com/scmn-dev/gh-api/pkg/cmd/factory"
	"github.com/scmn-dev/gh-api/pkg/iostreams"
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

type ColorScheme struct {
	IO *iostreams.IOStreams
}

func opts(f *cmdutil.Factory) ColorScheme {
	opts := ColorScheme{
		IO: f.IOStreams,
	}

	return opts
}

var cs = opts(factory.New()).IO.ColorScheme()

func AuthMessage() {
	fmt.Println("You're not authenticated, to authenticate run " + cs.Bold("secman auth login"))

	os.Exit(0)
}

func RunSMWin() string {
	return "run " + cs.Bold("sm-win start")
}

func PrintError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		fmt.Fprintf(out, "error connecting to %s\n", dnsError.Name)

		if debug {
			fmt.Fprintln(out, dnsError)
		}

		return
	}

	fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			fmt.Fprintln(out)
		}

		fmt.Fprintln(out, cmd.UsageString())
	}
}

// commands
func Auth(f *cmdutil.Factory) *cobra.Command {
	cmd := aCmd.NewCmdAuth(f)
	return cmd
}

func Config(f *cmdutil.Factory) *cobra.Command {
	cmd := cCmd.NewCmdConfig(f)
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
