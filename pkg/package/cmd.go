package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/scmn-dev/gh-api/pkg/cmdutil"
)

func PackageCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := cobra.Command{
		Use:   "package",
		Short: "Manage secman package",
		Long: heredoc.Docf(`
			Secman Package are repositories that provide more additional apps.
		`, "`"),
		Aliases: []string{"pkg", "packages"},
	}

	return &cmd
}
