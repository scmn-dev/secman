package root

import (
	"net/http"

	"github.com/MakeNowJust/heredoc"
	authCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/factory"
	repoCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/repo"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(f *cmdutil.Factory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secman <command> <subcommand> [flags]",
		Short: "Secman CLI",
		Long:  `Work seamlessly with GitHub from the command line.`,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			secman auth login
			secman repo clone scmn-dev/gh-api
		`),
		Annotations: map[string]string{
			"help:feedback": heredoc.Doc(`
				Open an issue using at https://github.com/scmn-dev/secman/pkg/api/gh/issues
			`),
		},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		rootHelpFunc(f, cmd, args)
	})
	cmd.SetUsageFunc(rootUsageFunc)
	cmd.SetFlagErrorFunc(rootFlagErrorFunc)

	// the `api` command should not inherit any extra HTTP headers
	bareHTTPCmdFactory := *f
	bareHTTPCmdFactory.HttpClient = bareHTTPClient(f, version)

	// below here at the commands that require the "intelligent" BaseRepo resolver
	repoResolvingCmdFactory := *f
	repoResolvingCmdFactory.BaseRepo = factory.SmartBaseRepoFunc(f)

	cmd.AddCommand(authCmd.NewCmdAuth(f))
	cmd.AddCommand(repoCmd.NewCmdRepo(&repoResolvingCmdFactory))

	// Help topics
	cmd.AddCommand(NewHelpTopic("environment"))
	cmd.AddCommand(NewHelpTopic("mintty"))
	referenceCmd := NewHelpTopic("reference")
	referenceCmd.SetHelpFunc(referenceHelpFn(f.IOStreams))
	cmd.AddCommand(referenceCmd)

	cmdutil.DisableAuthCheck(cmd)

	// this needs to appear last:
	referenceCmd.Long = referenceLong(cmd)
	return cmd
}

func bareHTTPClient(f *cmdutil.Factory, version string) func() (*http.Client, error) {
	return func() (*http.Client, error) {
		cfg, err := f.Config()
		if err != nil {
			return nil, err
		}

		return factory.NewHTTPClient(f.IOStreams, cfg, version, false)
	}
}
