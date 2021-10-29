package uni

import (
	"github.com/abdfnx/shell"
	commands "github.com/scmn-dev/secman-v1/tools/constants"
	"github.com/spf13/cobra"
)

type UninstallOptions struct {
	DeleteData bool
}

func Uninstall(runF func(*UninstallOptions) error) *cobra.Command {
	opts := UninstallOptions{}

	cmd := &cobra.Command{
		Use:   "uninstall",
		Aliases: []string{"un"},
		Short: "Uninstall Your Secman.",
		RunE: func(c *cobra.Command, args []string) error {
			if runF != nil {
				return runF(&opts)
			}

			return uniFunc(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.DeleteData, "delete-data", "d", false, "Delete .secman with secman.")

	return cmd
}

func uniFunc(opts *UninstallOptions) error {
	shell.ShellCmd(commands.Uninstall())

	if opts.DeleteData {
		shell.ShellCmd(commands.ClearData())
	}

	return nil
}
