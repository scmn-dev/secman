package app

import (
	"github.com/spf13/cobra"
	"github.com/scmn-dev/secman/pkg/initx"
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
