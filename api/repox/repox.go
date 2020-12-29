package repox

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cli/cli/git"
	"github.com/abdfnx/secman/v3/api/common"
)

type Interface interface {
	RepoName() string
	RepoOwner() string
	RepoHost() string
}

func New(owner, repo string) Interface {
	return NewWithHost(owner, repo, common.OverridableDefault())
}

func NewWithHost(owner, repo, hostname string) Interface {
	return &sRepo{
		owner:    owner,
		name:     repo,
		hostname: normalizeHostname(hostname),
	}
}

func FullName(r Interface) string {
	return fmt.Sprintf("%s/%s", r.RepoOwner(), r.RepoName())
}

func FromFullName(nwo string) (Interface, error) {
	if git.IsURL(nwo) {
		u, err := git.ParseURL(nwo)
		if err != nil {
			return nil, err
		}
		return FromURL(u)
	}

	parts := strings.SplitN(nwo, "/", 4)
	for _, p := range parts {
		if len(p) == 0 {
			return nil, fmt.Errorf(`expected the "[HOST/]OWNER/REPO" format, got %q`, nwo)
		}
	}
	switch len(parts) {
	case 3:
		return NewWithHost(parts[1], parts[2], normalizeHostname(parts[0])), nil
	case 2:
		return New(parts[0], parts[1]), nil
	default:
		return nil, fmt.Errorf(`expected the "[HOST/]OWNER/REPO" format, got %q`, nwo)
	}
}

func FromURL(u *url.URL) (Interface, error) {
	if u.Hostname() == "" {
		return nil, fmt.Errorf("no hostname detected")
	}

	parts := strings.SplitN(strings.Trim(u.Path, "/"), "/", 3)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid path: %s", u.Path)
	}

	return NewWithHost(parts[0], strings.TrimSuffix(parts[1], ".git"), u.Hostname()), nil
}

func normalizeHostname(h string) string {
	return strings.ToLower(strings.TrimPrefix(h, "www."))
}

func IsSame(a, b Interface) bool {
	return strings.EqualFold(a.RepoOwner(), b.RepoOwner()) &&
		strings.EqualFold(a.RepoName(), b.RepoName()) &&
		normalizeHostname(a.RepoHost()) == normalizeHostname(b.RepoHost())
}

func GenerateRepoURL(repo Interface, p string, args ...interface{}) string {
	baseURL := fmt.Sprintf("https://%s/%s/%s", repo.RepoHost(), repo.RepoOwner(), repo.RepoName())
	if p != "" {
		return baseURL + "/" + fmt.Sprintf(p, args...)
	}
	return baseURL
}

// TODO there is a parallel implementation for non-isolated commands
func FormatRemoteURL(repo Interface, protocol string) string {
	if protocol == "ssh" {
		return fmt.Sprintf("git@%s:%s/%s.git", repo.RepoHost(), repo.RepoOwner(), repo.RepoName())
	}

	return fmt.Sprintf("https://%s/%s/%s.git", repo.RepoHost(), repo.RepoOwner(), repo.RepoName())
}

type sRepo struct {
	owner    string
	name     string
	hostname string
}

func (r sRepo) RepoOwner() string {
	return r.owner
}

func (r sRepo) RepoName() string {
	return r.name
}

func (r sRepo) RepoHost() string {
	return r.hostname
}
