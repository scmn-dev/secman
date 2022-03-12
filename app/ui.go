package app

import (
	"github.com/spf13/cobra"
	"github.com/scmn-dev/secman/pkg/ui"
)

func UICMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ui",
		Short: "SMUI is your graphical secman.",
		Long: "SMUI is your graphical secman.",
	}

	cmd.AddCommand(UIOpenCMD)
	cmd.AddCommand(UIUpdateCMD)

	return cmd
}

var UIOpenCMD = &cobra.Command{
	Use:   "open",
	Aliases: []string{"."},
	Short: "Open Secman UI in default browser.",
	Long: "Open Secman UI in default browser.",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Open()
	},
}

var UIUpdateCMD = &cobra.Command{
	Use:   "update",
	Short: "Update SMUI to latest version.",
	Long: "Update SMUI to latest version.",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Update()
	},
}
