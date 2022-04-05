package cli

import (
	"github.com/scmn-dev/secman/v6/pkg/initx"
	"github.com/spf13/cobra"
)

var InitCMD = &cobra.Command{
	Use:   "init",
	Short: "Initialize ~/.secman .",
	Long:  "Initialize ~/.secman .",
	RunE: func(cmd *cobra.Command, args []string) error {
		initx.Init()

		return nil
	},
}
