package config

import (
	"fmt"
	"strings"

	"github.com/scmn-dev/secman/pkg/api/gh/core/config"
	cmdGet "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/cluster/get"
	cmdSet "github.com/scmn-dev/secman/pkg/api/gh/pkg/cmd/cluster/set"
	"github.com/scmn-dev/secman/pkg/api/gh/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdConfig(f *cmdutil.Factory) *cobra.Command {
	longDoc := strings.Builder{}
	longDoc.WriteString("Display or change configuration settings for secman.\n\n")
	longDoc.WriteString("Current respected settings:\n")
	for _, co := range config.ConfigOptions() {
		longDoc.WriteString(fmt.Sprintf("- %s: %s", co.Key, co.Description))
		if co.DefaultValue != "" {
			longDoc.WriteString(fmt.Sprintf(" (default: %q)", co.DefaultValue))
		}

		longDoc.WriteRune('\n')
	}

	cmd := &cobra.Command{
		Use:   "cluster <command>",
		Short: "Manage configuration of github for secman.",
		Long:  longDoc.String(),
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.AddCommand(cmdGet.NewCmdConfigGet(f, nil))
	cmd.AddCommand(cmdSet.NewCmdConfigSet(f, nil))

	return cmd
}
