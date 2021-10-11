package get

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scmn-dev/secman/pkg/api/gh/core/config"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmdutil"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/iostreams"
	"github.com/spf13/cobra"
)

type GetOptions struct {
	IO     *iostreams.IOStreams
	Config config.Config

	Hostname string
	Key      string
}

func NewCmdConfigGet(f *cmdutil.Factory, runF func(*GetOptions) error) *cobra.Command {
	opts := &GetOptions{
		IO: f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "get <key>",
		Short: "Print the value of a given configuration key",
		Example: heredoc.Doc(`
			secman config get git_protocol
			https
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := f.Config()
			if err != nil {
				return err
			}
			opts.Config = config
			opts.Key = args[0]

			if runF != nil {
				return runF(opts)
			}

			return getRun(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Hostname, "host", "", "", "Get per-host setting")

	return cmd
}

func getRun(opts *GetOptions) error {
	val, err := opts.Config.Get(opts.Hostname, opts.Key)
	if err != nil {
		return err
	}

	if val != "" {
		fmt.Fprintf(opts.IO.Out, "%s\n", val)
	}
	return nil
}
