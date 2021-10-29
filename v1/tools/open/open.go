package openx

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"encoding/base64"

	"github.com/scmn-dev/secman-v1/pkg/api/gh/api"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/core/ghrepo"
)

var NotFoundError = errors.New("not found")

func IsMarkdownFile(filename string) bool {
	// kind of gross, but i'm assuming that 90% of the time the suffix will just be .md. it didn't
	// seem worth executing a regex for this given that assumption.
	return strings.HasSuffix(filename, ".md") ||
		strings.HasSuffix(filename, ".markdown") ||
		strings.HasSuffix(filename, ".mdown") ||
		strings.HasSuffix(filename, ".mkdown")
}

type RepoReadme struct {
	Filename string
	Content  string
	BaseURL  string
}

func RepositoryReadme(client *http.Client, repo ghrepo.Interface, branch string) (*RepoReadme, error) {
	apiClient := api.NewClientFromHTTP(client)
	var response struct {
		Name    string
		Content string
		HTMLURL string `json:"html_url"`
	}

	err := apiClient.REST(repo.RepoHost(), "GET", GetReadmePath(repo, branch), nil, &response)
	if err != nil {
		var httpError api.HTTPError
		if errors.As(err, &httpError) && httpError.StatusCode == 404 {
			return nil, NotFoundError
		}
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(response.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode readme: %w", err)
	}

	return &RepoReadme{
		Filename: response.Name,
		Content:  string(decoded),
		BaseURL:  response.HTMLURL,
	}, nil
}

func GetReadmePath(repo ghrepo.Interface, branch string) string {
	path := fmt.Sprintf("repos/%s/readme", ghrepo.FullName(repo))
	if branch != "" {
		path = fmt.Sprintf("%s?ref=%s", path, branch)
	}
	return path
}


func GenerateBranchURL(r ghrepo.Interface, branch string) string {
	if branch == "" {
		return ghrepo.GenerateRepoURL(r, "")
	}

	return ghrepo.GenerateRepoURL(r, "tree/%s", url.QueryEscape(branch))
}
