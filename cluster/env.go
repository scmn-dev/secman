package cluster

import (
	"fmt"
	"os"

	"github.com/scmn-dev/gh-api/core/ghinstance"
)

const (
	SM_GH_HOST                 = "SM_GH_HOST"
	SM_GH_TOKEN                = "SM_GH_TOKEN"
	SM_GITHUB_TOKEN            = "SM_GITHUB_TOKEN"
	SM_GH_ENTERPRISE_TOKEN     = "SM_GH_ENTERPRISE_TOKEN"
	SM_GITHUB_ENTERPRISE_TOKEN = "SM_GITHUB_ENTERPRISE_TOKEN"
)

type ReadOnlyEnvError struct {
	Variable string
}

func (e *ReadOnlyEnvError) Error() string {
	return fmt.Sprintf("read-only value in %s", e.Variable)
}

func InheritEnv(c Cluster) Cluster {
	return &envCluster{Cluster: c}
}

type envCluster struct {
	Cluster
}

func (c *envCluster) Hosts() ([]string, error) {
	hasDefault := false
	hosts, err := c.Cluster.Hosts()
	for _, h := range hosts {
		if h == ghinstance.Default() {
			hasDefault = true
		}
	}

	token, _ := AuthTokenFromEnv(ghinstance.Default())
	if (err != nil || !hasDefault) && token != "" {
		hosts = append([]string{ghinstance.Default()}, hosts...)
		return hosts, nil
	}

	return hosts, err
}

func (c *envCluster) DefaultHost() (string, error) {
	val, _, err := c.DefaultHostWithSource()
	return val, err
}

func (c *envCluster) DefaultHostWithSource() (string, string, error) {
	if host := os.Getenv(SM_GH_HOST); host != "" {
		return host, SM_GH_HOST, nil
	}

	return c.Cluster.DefaultHostWithSource()
}

func (c *envCluster) Get(hostname, key string) (string, error) {
	val, _, err := c.GetWithSource(hostname, key)
	return val, err
}

func (c *envCluster) GetWithSource(hostname, key string) (string, string, error) {
	if hostname != "" && key == "oauth_token" {
		if token, env := AuthTokenFromEnv(hostname); token != "" {
			return token, env, nil
		}
	}

	return c.Cluster.GetWithSource(hostname, key)
}

func (c *envCluster) CheckWriteable(hostname, key string) error {
	if hostname != "" && key == "oauth_token" {
		if token, env := AuthTokenFromEnv(hostname); token != "" {
			return &ReadOnlyEnvError{Variable: env}
		}
	}

	return c.Cluster.CheckWriteable(hostname, key)
}

func AuthTokenFromEnv(hostname string) (string, string) {
	if ghinstance.IsEnterprise(hostname) {
		if token := os.Getenv(SM_GH_ENTERPRISE_TOKEN); token != "" {
			return token, SM_GH_ENTERPRISE_TOKEN
		}

		return os.Getenv(SM_GITHUB_ENTERPRISE_TOKEN), SM_GITHUB_ENTERPRISE_TOKEN
	}

	if token := os.Getenv(SM_GH_TOKEN); token != "" {
		return token, SM_GH_TOKEN
	}

	return os.Getenv(SM_GITHUB_TOKEN), SM_GITHUB_TOKEN
}

func AuthTokenProvidedFromEnv() bool {
	return os.Getenv(SM_GH_ENTERPRISE_TOKEN) != "" ||
		os.Getenv(SM_GITHUB_ENTERPRISE_TOKEN) != "" ||
		os.Getenv(SM_GH_TOKEN) != "" ||
		os.Getenv(SM_GITHUB_TOKEN) != ""
}

func IsHostEnv(src string) bool {
	return src == SM_GH_HOST
}

func IsEnterpriseEnv(src string) bool {
	return src == SM_GH_ENTERPRISE_TOKEN || src == SM_GITHUB_ENTERPRISE_TOKEN
}
