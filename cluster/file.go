package cluster

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type fileCluster struct {
	ClusterMap
	documentRoot *yaml.Node
}

type HostCluster struct {
	ClusterMap
	Host string
}

func (c *fileCluster) Root() *yaml.Node {
	return c.ClusterMap.Root
}

func (c *fileCluster) Get(hostname, key string) (string, error) {
	val, _, err := c.GetWithSource(hostname, key)
	return val, err
}

func (c *fileCluster) GetWithSource(hostname, key string) (string, string, error) {
	if hostname != "" {
		var notFound *NotFoundError

		hostClr, err := c.clusterForHost(hostname)
		if err != nil && !errors.As(err, &notFound) {
			return "", "", err
		}

		var hostValue string
		if hostClr != nil {
			hostValue, err = hostClr.GetStringValue(key)
			if err != nil && !errors.As(err, &notFound) {
				return "", "", err
			}
		}

		if hostValue != "" {
			return hostValue, HostsClusterFile(), nil
		}
	}

	defaultSource := ClusterFile()

	value, err := c.GetStringValue(key)

	var notFound *NotFoundError

	if err != nil && errors.As(err, &notFound) {
		return defaultFor(key), defaultSource, nil
	} else if err != nil {
		return "", defaultSource, err
	}

	if value == "" {
		return defaultFor(key), defaultSource, nil
	}

	return value, defaultSource, nil
}

func (c *fileCluster) clusterForHost(hostname string) (*HostCluster, error) {
	hosts, err := c.hostEntries()
	if err != nil {
		return nil, err
	}

	for _, hc := range hosts {
		if strings.EqualFold(hc.Host, hostname) {
			return hc, nil
		}
	}

	return nil, &NotFoundError{fmt.Errorf("could not find cluster entry for %q", hostname)}
}

func defaultFor(key string) string {
	for _, co := range clusterOptions {
		if co.Key == key {
			return co.DefaultValue
		}
	}

	return ""
}

func (c *fileCluster) hostEntries() ([]*HostCluster, error) {
	entry, err := c.FindEntry("hosts")
	if err != nil {
		return []*HostCluster{}, nil
	}

	hostClusters, err := c.parseHosts(entry.ValueNode)
	if err != nil {
		return nil, fmt.Errorf("could not parse hosts cluster: %w", err)
	}

	return hostClusters, nil
}

func (c *fileCluster) parseHosts(hostsEntry *yaml.Node) ([]*HostCluster, error) {
	hostClusters := []*HostCluster{}

	for i := 0; i < len(hostsEntry.Content)-1; i = i + 2 {
		hostname := hostsEntry.Content[i].Value
		hostRoot := hostsEntry.Content[i+1]
		hostCluster := HostCluster{
			ClusterMap: ClusterMap{Root: hostRoot},
			Host:      hostname,
		}

		hostClusters = append(hostClusters, &hostCluster)
	}

	if len(hostClusters) == 0 {
		return nil, errors.New("could not find any host clusters")
	}

	return hostClusters, nil
}
