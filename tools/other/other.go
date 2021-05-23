package other

import (
	"fmt"
	"io"
	"errors"
	"net"
	"strings"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func PrintError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		fmt.Fprintf(out, "error connecting to %s\n", dnsError.Name)

		if debug {
			fmt.Fprintln(out, dnsError)
		}

		return
	}

	fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			fmt.Fprintln(out)
		}

		fmt.Fprintln(out, cmd.UsageString())
	}
}
