package secman

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scmn-dev/secman/v6/cli"
	"github.com/scmn-dev/secman/v6/cli/factory"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/pkg/options"
	"github.com/spf13/cobra"
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
		cli.AuthCMD(),
		cli.DocsCMD,
		cli.EncryptCMD(),
		cli.DoctorCMD(version),
		cli.InitCMD,
		cli.InfoCMD(version),
		cli.InsertCMD(),
		cli.GenerateCMD(),
		cli.FilesCMD(),
		cli.ReadCMD(),
		cli.EditCMD(),
		cli.ListCMD(),
		cli.DeleteCMD(),
		cli.UICMD(),
		cli.WhoamiCMD(),
		versionCmd,
	)

	return rootCmd
}
