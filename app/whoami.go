package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/internal/config"
)

func WhoamiCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Display the current user.",
		Long:  "Display the current user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := lipgloss.NewStyle().PaddingLeft(2)

			primary := lipgloss.NewStyle().Foreground(lipgloss.Color(constants.PRIMARY_COLOR))
			bold := lipgloss.NewStyle().Bold(true)

			fmt.Println(s.Render("\n👊 Hi ") + primary.Render(config.Config("config.name")) + " <" + bold.Render(config.Config("config.user")) + ">")

			return nil
		},
	}

	return cmd
}
