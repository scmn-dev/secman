package open

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"syscall"
	"text/template"

	"github.com/MakeNowJust/heredoc"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/api"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/core/ghinstance"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/core/ghrepo"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/cmdutil"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/iostreams"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/markdown"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/utils"
	"github.com/spf13/cobra"
	openx "github.com/scmn-dev/secman-v1/tools/open"
	git_config "github.com/scmn-dev/secman-v1/tools/config"
)

type browser interface {
	Browse(string) error
}

type OpenOptions struct {
	HttpClient func() (*http.Client, error)
	IO         *iostreams.IOStreams
	BaseRepo   func() (ghrepo.Interface, error)
	Browser    browser

	RepoArg string
	Web     bool
	Branch  string
}

var NotFoundError = errors.New("not found")

func Open(f *cmdutil.Factory, runF func(*OpenOptions) error) *cobra.Command {
	opts := OpenOptions{
		IO:         f.IOStreams,
		HttpClient: f.HttpClient,
		BaseRepo:   f.BaseRepo,
		Browser:    f.Browser,
	}

	cmd := &cobra.Command{
		Use:   "open",
		Short: OpenHelp(),
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.RepoArg = args[0]
			}

			if runF != nil {
				return runF(&opts)
			}

			return openRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Web, "web", "w", false, "Open your repo in the browser")

	return cmd
}

func openRun(opts *OpenOptions) error {
	httpClient, err := opts.HttpClient()
	if err != nil {
		return err
	}

	var toOpen ghrepo.Interface
	apiClient := api.NewClientFromHTTP(httpClient)
	openURL := ".secman"
	if !strings.Contains(openURL, "/") {
		currentUser, err := api.CurrentLoginName(apiClient, ghinstance.Default())
		if err != nil {
			return err
		}

		openURL = currentUser + "/" + openURL
	}

	toOpen, err = ghrepo.FromFullName(openURL)

	if err != nil {
		return fmt.Errorf("argument error: %w", err)
	}

	repo, err := api.GitHubRepo(apiClient, toOpen)
	if err != nil {
		return err
	}

	if opts.Web {
		openURL := openx.GenerateBranchURL(toOpen, opts.Branch)
		if opts.IO.IsStdoutTTY() {
			fmt.Fprintf(opts.IO.ErrOut, "Opening %s in your browser.\n", utils.DisplayURL(openURL))
		}

		return opts.Browser.Browse(openURL)
	}

	fullName := ghrepo.FullName(toOpen)

	readme, err := openx.RepositoryReadme(httpClient, toOpen, opts.Branch)
	if err != nil && err != NotFoundError {
		return err
	}

	if err != nil && err != NotFoundError {
		return err
	}

	opts.IO.DetectTerminalTheme()

	err = opts.IO.StartPager()
	if err != nil {
		return err
	}
	defer opts.IO.StopPager()

	stdout := opts.IO.Out

	if !opts.IO.IsStdoutTTY() {
		fmt.Fprintf(stdout, "name:\t%s\n", fullName)
		fmt.Fprintf(stdout, "description:\t%s\n", repo.Description)
		return nil
	}

	repoTmpl := heredoc.Doc(`
		{{.FullName}}
		{{.Description}}

		{{.Readme}}

		{{.Open}}
	`)

	tmpl, err := template.New("repo").Parse(repoTmpl)
	if err != nil {
		return err
	}

	cs := opts.IO.ColorScheme()

	var readmeContent string
	if readme == nil {
		readmeContent = cs.Gray("This repository does not have a README")
	} else if openx.IsMarkdownFile(readme.Filename) {
		var err error
		style := markdown.GetStyle(opts.IO.TerminalTheme())
		readmeContent, err = markdown.RenderWithBaseURL(readme.Content, style, readme.BaseURL)
		if err != nil {
			return fmt.Errorf("error rendering markdown: %w", err)
		}
	} else {
		readmeContent = readme.Content
	}

	description := repo.Description
	if description == "" {
		description = cs.Gray("No description provided")
	}

	repoData := struct {
		FullName    string
		Description string
		Readme      string
		Open        string
	}{
		FullName:    cs.Bold(fullName),
		Description: description,
		Readme:      readmeContent,
		Open:        cs.Gray(fmt.Sprintf("Open this repository on GitHub: %s", openURL)),
	}

	err = tmpl.Execute(stdout, repoData)
	if err != nil && !errors.Is(err, syscall.EPIPE) {
		return err
	}

	return nil
}

func OpenHelp() string {
	return git_config.GitConfigWithMsg("Open Your Private Repo (", "/.secman).")
}
