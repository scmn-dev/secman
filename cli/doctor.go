package cli

import (
	"github.com/scmn-dev/secman/v6/pkg/doctor"
	"github.com/spf13/cobra"
)

func DoctorCMD(buildVersion string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Show the status of secman.",
		Long: "Show the status of secman.",
		Aliases: []string{"check"},
		Run: func(cmd *cobra.Command, args []string) {
			doctor.Doctor(buildVersion)
		},
	}

	cmd.AddCommand(FixCMD(buildVersion))

	return cmd
}

func FixCMD(buildVersion string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "fix",
		Short: "Show Information about how to fix secman issues and bugs.",
		Long: "Show Information about how to fix secman issues and bugs.",
		Run: func(cmd *cobra.Command, args []string) {
			doctor.Fix(buildVersion)
		},
	}

	return cmd
}
