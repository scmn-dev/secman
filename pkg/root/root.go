package root

import (
	"fmt"
	"strconv"

	"github.com/MakeNowJust/heredoc"
	"github.com/secman-team/gh-api/pkg/cmd/factory"
	"github.com/secman-team/gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
	sync "github.com/secman-team/secman/pkg/api"
	"github.com/secman-team/secman/pkg/clean"
	"github.com/secman-team/secman/pkg/edit"
	"github.com/secman-team/secman/pkg/fetch"
	"github.com/secman-team/secman/pkg/open"
	"github.com/secman-team/secman/pkg/gen"
	"github.com/secman-team/secman/pkg/initialize"
	"github.com/secman-team/secman/pkg/insert"
	"github.com/secman-team/secman/pkg/pio"
	"github.com/secman-team/secman/pkg/show"
	upg "github.com/secman-team/secman/pkg/upgrade"
	uni "github.com/secman-team/secman/pkg/uninstall"
	authx "github.com/secman-team/secman/tools/auth"
	configx "github.com/secman-team/secman/tools/config"
	repox "github.com/secman-team/secman/tools/repo"
	checker "github.com/secman-team/version-checker"
)

var (
	copyPass bool

	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Clean your ~/.secman (delete it).",
		Run: func(cmd *cobra.Command, args []string) {
			clean.Clean()
			checker.Checker()
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize your secman vault.",
		Long:  "Initialize the .secman directory, and generate your secret keys.",
		Run: func(cmd *cobra.Command, args []string) {
			initialize.Init()
			checker.Checker()
		},
	}

	upgradeCmd = &cobra.Command{
		Use:     "upgrade",
		Aliases: []string{"upg"},
		Short:   "Upgrade your secman if there's a new release.",
		Run: func(cmd *cobra.Command, args []string) {
			upg.Upgrade()
		},
	}

	insertCmd = &cobra.Command{
		Use:     "insert",
		Short:   "Insert a file or password in to your vault.",
		Example: "secman insert core/docker.com",
		Args:    cobra.RangeArgs(1, 2),
		Long: `Add a site to your password store. This site can optionally be a part
of a group by prepending a group name and slash to the site name.
Will prompt for confirmation when a site path is not unique.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 2 {
				path := args[0]
				filename := args[1]
				insert.File(path, filename)
			} else {
				pathName := args[0]
				insert.Password(pathName)
			}

			sync.PushSync()
			checker.Checker()
		},
	}

	showCmd = &cobra.Command{
		Use:     "show",
		Aliases: []string{"read"},
		Example: "secman show core/docker-password",
		Short:   "Print the password of a secman entry.",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			show.Site(path, copyPass)
			checker.Checker()
		},
	}

	generateCmd = &cobra.Command{
		Use:     "gen",
		Aliases: []string{"generate"},
		Short:   "Generate a secure password.",
		Example: "secman generate",
		Long: `Prints a randomly generated password. The length of this password defaults
to (24). If a password length is specified as greater than 2048 then generate
will fail.`,
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			pwlen := -1
			if len(args) != 0 {
				pwlenStr := args[0]
				pwlenint, err := strconv.Atoi(pwlenStr)
				if err != nil {
					pwlen = -1
				} else {
					pwlen = pwlenint
				}
			}

			pass := gen.Generate(pwlen)
			fmt.Println(pass)
			checker.Checker()
		},
	}

	findCmd = &cobra.Command{
		Use:     "find",
		Aliases: []string{"ls"},
		Example: "secman find code.com",
		Short:   "Find a site that contains the site-path.",
		Long: `Prints all sites that contain the site-path. Used to print just
one group or all sites that contain a certain word in the group or name.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			show.Find(path)
			checker.Checker()
		},
	}

	renameCmd = &cobra.Command{
		Use:     "rename",
		Short:   "Rename an entry in the password vault.",
		Example: "secman rename core/docker.com",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.Rename(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	editCmd = &cobra.Command{
		Use:     "edit",
		Aliases: []string{"update"},
		Short:   "Change the password of a site in the vault.",
		Example: "secman edit core/docker.com",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.Edit(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	removeCmd = &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Example: "secman remove core/docker.com",
		Short:   "Remove a site from the password vault by specifying the entire site-path.",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			edit.RemovePassword(path)
			sync.PushSync()
			checker.Checker()
		},
	}

	fetchCmd = &cobra.Command{
		Use:     "fetch",
		Example: "secman fetch",
		Short:   "Fetch if there is a new password/s in ~/.secman.",
		Run: func(cmd *cobra.Command, args []string) {
			fetch.FetchSECDIR()
			checker.Checker()
		},
	}

	syncCmd = sync.Sync()
	openCmd = open.Open(factory.New("x"), nil)
	uninstallCmd  = uni.Uninstall(nil)

	// with github
	repoCmd = repox.Repo(factory.New("x"))
	authCmd = authx.Auth(factory.New("x"))
	configCmd = configx.Config(factory.New("x"))
)

func NewCmdRoot(f *cmdutil.Factory, version string, versionDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secman <command> <subcommand> [flags]",
		Short: "Secman CLI",
		Long:  `Be 🔒, With Secman.`,

		SilenceErrors: true,
		Example: heredoc.Doc(`
			secman insert new-password
			secman gen 15
			secman sync start
			secman find new-password 
		`),
		Annotations: map[string]string{
			"help:feedback": heredoc.Doc(`
				Open an issue at https://github.com/secman-team/secman/issues
			`),
		},
		Run: func(cmd *cobra.Command, args []string) {
			if exists, _ := pio.PassFileDirExists(); exists {
				show.ListAll()
			} else {
				cmd.Help()
			}

			checker.Checker()
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Aliases: []string{"ver"},
		Short: "Print the version of your secman binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("secman " + version + " " + versionDate)
			fmt.Println("https://github.com/secman-team/secman/releases/tag/" + version)
			checker.Checker()
		},
	}

	verxCmd := &cobra.Command{
		Use: "verx",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(version)
		},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	cmd.PersistentFlags().Bool("help", false, "Help for secman")
	cmd.SetHelpFunc(helpHelper)
	cmd.SetUsageFunc(rootUsageFunc)
	cmd.SetFlagErrorFunc(rootFlagErrorFunc)

	cmd.AddCommand(authCmd)
	cmd.AddCommand(repoCmd)
	cmd.AddCommand(configCmd)
	cmd.AddCommand(cleanCmd)
	cmd.AddCommand(fetchCmd)
	cmd.AddCommand(findCmd)
	cmd.AddCommand(generateCmd)
	cmd.AddCommand(initCmd)
	cmd.AddCommand(insertCmd)
	cmd.AddCommand(removeCmd)
	cmd.AddCommand(editCmd)
	cmd.AddCommand(renameCmd)
	cmd.AddCommand(showCmd)
	cmd.AddCommand(versionCmd)
	cmd.AddCommand(verxCmd)
	cmd.AddCommand(upgradeCmd)
	cmd.AddCommand(uninstallCmd)
	cmd.AddCommand(syncCmd)
	cmd.AddCommand(openCmd)

	cmdutil.DisableAuthCheck(cmd)

	return cmd
}
