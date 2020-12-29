package alias

import (
	"github.com/MakeNowJust/heredoc"
	deleteCmd "github.com/abdfnx/secman/v3/api/cmd/alias/delete"
	listCmd "github.com/abdfnx/secman/v3/api/cmd/alias/list"
	setCmd "github.com/abdfnx/secman/v3/api/cmd/alias/set"
	"github.com/abdfnx/secman/v3/api/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAlias(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias <command>",
		Short: "Create command shortcuts",
		Long: heredoc.Doc(`
			Aliases can be used to make shortcuts for sm commands or to compose multiple commands.

			Run "sm help alias set" to learn more.
		`),
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(deleteCmd.NewCmdDelete(f, nil))
	cmd.AddCommand(listCmd.NewCmdList(f, nil))
	cmd.AddCommand(setCmd.NewCmdSet(f, nil))

	return cmd
}
