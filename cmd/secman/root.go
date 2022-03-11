package secman

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
	"github.com/scmn-dev/secman/app"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/cmd/factory"
	"github.com/scmn-dev/secman/pkg/options"
)

var opts = options.RootOptions{
	Version: false,
}

func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	const desc = `👊 Human-friendly and amazing TUI secrets manager.`

	var rootCmd = &cobra.Command{
		Use:   "secman <subcommand> [flags]",
		Short:  desc,
		Long: desc,
		SilenceErrors: true,
		Example: constants.HelpExamples,
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/scmn-dev/secman/issues
			`),
		},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Version {
				fmt.Println("secman version " + version + " " + buildDate)
			} else {
				cmd.Help()
			}
		},
	}

	rootCmd.SetOut(f.IOStreams.Out)
	rootCmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	rootCmd.PersistentFlags().Bool("help", false, "Help for secman")
	rootCmd.SetHelpFunc(helpHelper)
	rootCmd.SetUsageFunc(rootUsageFunc)
	rootCmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "Print the version of your secman binary.")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of your secman binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("secman version " + version + " " + buildDate)
		},
	}

	// add secman commands to root command
	rootCmd.AddCommand(
		app.AuthCMD(),
		app.DocsCMD,
		app.DoctorCMD(version),
		app.InitCMD,
		app.InfoCMD(version),
		app.InsertCMD(),
		app.GenerateCMD(),
		app.ReadCMD(),
		app.EditCMD(),
		app.ListCMD(),
		app.DeleteCMD(),
		app.WhoamiCMD(),
		versionCmd,
	)

	return rootCmd
}
