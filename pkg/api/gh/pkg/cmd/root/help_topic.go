package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var HelpTopics = map[string]map[string]string{
	"mintty": {
		"short": "Information about using secman and gh-api with MinTTY",
		"long": heredoc.Doc(`
			MinTTY is the terminal emulator that comes by default with Git
			for Windows.  It has known issues with secman's ability to prompt a
			user for input.
			There are a few workarounds to make secman work with MinTTY:
			- Reinstall Git for Windows, checking "Enable experimental support for pseudo consoles".

			- Use a different terminal emulator with Git for Windows like Windows Terminal.
			  You can run "C:\Program Files\Git\bin\bash.exe" from any terminal emulator to continue
			  using all of the tooling in Git For Windows without MinTTY.

			- Prefix invocations of secman with winpty, eg: "winpty secman auth login".
			  NOTE: this can lead to some UI bugs.
		`),
	},
	"environment": {
		"short": "Environment variables that can be used with secman",
		"long": heredoc.Doc(`
			GH_TOKEN, GITHUB_TOKEN (in order of precedence): an authentication token for github.com
			API requests. Setting this avoids being prompted to authenticate and takes precedence over
			previously stored credentials.

			GH_ENTERPRISE_TOKEN, GITHUB_ENTERPRISE_TOKEN (in order of precedence): an authentication
			token for API requests to GitHub Enterprise. When setting this, also set GH_HOST.
			GH_HOST: specify the GitHub hostname for commands that would otherwise assume the
			"github.com" host when not in a context of an existing repository.

			GH_REPO: specify the GitHub repository in the "[HOST/]OWNER/REPO" format for commands
			that otherwise operate on a local repository.

			GH_EDITOR, GIT_EDITOR, VISUAL, EDITOR (in order of precedence): the editor tool to use
			for authoring text.

			BROWSER: the web browser to use for opening links.

			DEBUG: set to any value to enable verbose output to standard error. Include values "api"
			or "oauth" to print detailed information about HTTP requests or authentication flow.

			GH_PAGER, PAGER (in order of precedence): a terminal paging program to send standard output
			to, e.g. "less".

			GLAMOUR_STYLE: the style to use for rendering Markdown. See
			https://github.com/charmbracelet/glamour#styles

			NO_COLOR: set to any value to avoid printing ANSI escape sequences for color output.

			CLICOLOR: set to "0" to disable printing ANSI colors in output.

			CLICOLOR_FORCE: set to a value other than "0" to keep ANSI colors in output
			even when the output is piped.

			GH_FORCE_TTY: set to any value to force terminal-style output even when the output is
			redirected. When the value is a number, it is interpreted as the number of columns
			available in the viewport. When the value is a percentage, it will be applied against
			the number of columns available in the current viewport.

			GH_NO_UPDATE_NOTIFIER: set to any value to disable update notifications. By default, secman
			checks for new releases once every 24 hours and displays an upgrade notice on standard
			error if a newer version was found.

			GH_CONFIG_DIR: the directory where gh will store configuration files. Default:
			"$XDG_CONFIG_HOME/gh" or "$HOME/.config/gh".
		`),
	},
	"reference": {
		"short": "A comprehensive reference of all secman commands",
	},
}

func NewHelpTopic(topic string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    topic,
		Short:  HelpTopics[topic]["short"],
		Long:   HelpTopics[topic]["long"],
		Hidden: true,
		Annotations: map[string]string{
			"markdown:generate": "true",
			"markdown:basename": "gh_help_" + topic,
		},
	}

	cmd.SetHelpFunc(helpTopicHelpFunc)
	cmd.SetUsageFunc(helpTopicUsageFunc)

	return cmd
}

func helpTopicHelpFunc(command *cobra.Command, args []string) {
	command.Print(command.Long)
}

func helpTopicUsageFunc(command *cobra.Command) error {
	command.Printf("Usage: gh help %s", command.Use)
	return nil
}
