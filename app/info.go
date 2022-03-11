package app

import (
	"github.com/spf13/cobra"
	"github.com/scmn-dev/secman/pkg/info"
)

func InfoCMD(version string) *cobra.Command{
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Information about the secman CLI.",
		Long:  "Information about the secman CLI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			info.Info(version)

			return nil
		},
	}

	return cmd
}
