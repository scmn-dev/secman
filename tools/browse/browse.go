package browsex

import (
	bCmd "github.com/secman-team/gh-api/pkg/cmd/browse"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func Browse(f *cmdutil.Factory) *cobra.Command {
	cmd := bCmd.NewCmdBrowse(f, nil)
	return cmd
}
