package app

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/abdfnx/gosh"
	"github.com/spf13/cobra"
	"github.com/scmn-dev/browser"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/pkg/auth/login"
	"github.com/scmn-dev/secman/internal/shared"
	"github.com/scmn-dev/secman/pkg/auth/logout"
	"github.com/scmn-dev/secman/pkg/auth/refresh"
)

func AuthCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage secman's authentication state.",
		Long: "Manage secman's authentication state.",
	}

	cmd.AddCommand(CreateCMD)
	cmd.AddCommand(LoginCMD())
	cmd.AddCommand(LogoutCMD())
	cmd.AddCommand(RefreshCMD)

	return cmd
}

func LoginCMD() *cobra.Command{
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with Secman.",
		Long: "Authenticate with Secman.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if AuthOpts.Username != "" && AuthOpts.PasswordStdin {
				bytes, err := ioutil.ReadAll(os.Stdin)

				if err != nil {
					return err
				}

				AuthOpts.Password = string(bytes)

				if AuthOpts.Password == "" {
					return fmt.Errorf("password cannot be empty")
				}

				err, out, errout := gosh.RunOutput("scc auth -e " + AuthOpts.Username + " -m " + AuthOpts.Password)

				if err != nil {
					fmt.Println(errout)
					fmt.Println(err)
					return err
				}

				st := shared.DefaultStyles()

				if strings.Contains(out, "200") {
					fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Auth") + "\n\n" + st.Success.Render("🎉 Welcome Back")).String())
				} else {
					fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Auth") + "\n\n" + st.Error.Render("🚫 Login Failed")).String())
				}
			} else {
				if err := tea.NewProgram(login.Login()).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(1)
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&AuthOpts.Username, "username", "u", "", "Username to login with.")
	cmd.Flags().BoolVarP(&AuthOpts.PasswordStdin, "password", "p", false, "Read password from standard input")

	return cmd
}

func LogoutCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of the current user account.",
		Long: "Logout of the current user account.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if AuthOpts.ConfirmLogout {
				gosh.Run("scc logout")
				msg := "✔ Logged out successfully"

				logo := "\n" + lipgloss.NewStyle().
					Foreground(lipgloss.Color("#fff")).
					Background(lipgloss.Color(constants.PRIMARY_COLOR)).
					Padding(0, 1).
					SetString("Secman Auth").
					String() + "\n\n"

				fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(logo + msg).String())
			} else {
				if err := tea.NewProgram(logout.Logout()).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(1)
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&AuthOpts.ConfirmLogout, "yes", "y", false, "Logout without confirmation.")

	return cmd
}

var CreateCMD = &cobra.Command{
	Use:   "create",
	Short: "Create a new secman account.",
	Long: "Create a new secman account at https://auth.secman.dev",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := browser.OpenURL("https://auth.secman.dev")

		if err != nil {
			fmt.Printf("could not open browser: %s\n", err)

			os.Exit(1)
		}

		return nil
	},
}

var RefreshCMD = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh the current user account.",
	Long: "Refresh the current user account. This will invalidate the current access token.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := tea.NewProgram(refresh.Refresh()).Start(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
		
		return nil
	},
}
