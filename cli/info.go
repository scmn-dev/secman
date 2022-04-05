package cli

import (
	"github.com/scmn-dev/secman/v6/pkg/info"
	"github.com/spf13/cobra"
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
