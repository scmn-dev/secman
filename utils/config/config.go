package configx

import (
	configCmd "github.com/secman-team/gh-api/pkg/cmd/config"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func Config(f *cmdutil.Factory) *cobra.Command {
	cmd := configCmd.NewCmdConfig(f)

	return cmd
}
