package configx

import (
	"fmt"
	"log"
	configCmd "github.com/secman-team/gh-api/pkg/cmd/config"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
	"strings"
	"github.com/secman-team/shell"
)

func Config(f *cmdutil.Factory) *cobra.Command {
	cmd := configCmd.NewCmdConfig(f)

	return cmd
}

func GitConfig() string {
	cmd := "git config user.name"

	err, username, errout := shell.SHCoreOut(cmd, cmd)

	uname := strings.TrimSuffix(username, "\n")

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	if uname != "" {
		return uname
	} else {
		return ":USERNAME"
	}
}
