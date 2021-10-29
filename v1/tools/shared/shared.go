package shared

import (
	"os"
	"fmt"
	"io"
	"errors"
	"net"
	"strings"

	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/cmd/factory"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/iostreams"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/cmdutil"

	"github.com/spf13/cobra"
)

type ColorScheme struct {
	IO *iostreams.IOStreams
}

func opts(f *cmdutil.Factory) ColorScheme {
	opts := ColorScheme{
		IO: f.IOStreams,
	}

	return opts
}

var cs = opts(factory.New()).IO.ColorScheme()

func AuthMessage() {
	fmt.Println("You're not authenticated, to authenticate run " + cs.Bold("secman auth login"))

	os.Exit(0)
}

func RunSMWin() string {
	return "run " + cs.Bold("sm-win start")
}

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
