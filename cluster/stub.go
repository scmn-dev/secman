package cluster

import (
	"errors"
)

type ClusterStub map[string]string

func genKey(host, key string) string {
	if host != "" {
		return host + ":" + key
	}

	return key
}

func (c ClusterStub) Get(host, key string) (string, error) {
	val, _, err := c.GetWithSource(host, key)
	return val, err
}

func (c ClusterStub) GetWithSource(host, key string) (string, string, error) {
	if v, found := c[genKey(host, key)]; found {
		return v, "(memory)", nil
	}

	return "", "", errors.New("not found")
}

func (c ClusterStub) Set(host, key, value string) error {
	c[genKey(host, key)] = value
	return nil
}

func (c ClusterStub) Hosts() ([]string, error) {
	return nil, nil
}

func (c ClusterStub) UnsetHost(hostname string) {}

func (c ClusterStub) CheckWriteable(host, key string) error {
	return nil
}

func (c ClusterStub) Write() error {
	c["_written"] = "true"
	return nil
}

func (c ClusterStub) DefaultHost() (string, error) {
	return "", nil
}

func (c ClusterStub) DefaultHostWithSource() (string, string, error) {
	return "", "", nil
}
