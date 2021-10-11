package auth

import (
	gitCredentialCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/gitcredential"
	authGetUsernameCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/get-username"
	authLoginCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/login"
	authLogoutCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/logout"
	authRefreshCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/refresh"
	authStatusCmd "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/auth/status"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAuth(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Login, logout, and refresh your authentication with github.",
		Long:  `Manage secman's authentication state.`,
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(authGetUsernameCmd.GetUsername())
	cmd.AddCommand(authLoginCmd.NewCmdLogin(f, nil))
	cmd.AddCommand(authLogoutCmd.NewCmdLogout(f, nil))
	cmd.AddCommand(authStatusCmd.NewCmdStatus(f, nil))
	cmd.AddCommand(authRefreshCmd.NewCmdRefresh(f, nil))
	cmd.AddCommand(gitCredentialCmd.NewCmdCredential(f, nil))

	return cmd
}
