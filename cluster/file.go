package cluster

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/scmn-dev/gh-api/core/ghinstance"
	yaml "gopkg.in/yaml.v3"
)

// This type implements a Cluster interface and represents a cluster file on disk.
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

		hostCls, err := c.clusterForHost(hostname)
		if err != nil && !errors.As(err, &notFound) {
			return "", "", err
		}

		var hostValue string
		if hostCls != nil {
			hostValue, err = hostCls.GetStringValue(key)
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

func (c *fileCluster) Set(hostname, key, value string) error {
	if hostname == "" {
		return c.SetStringValue(key, value)
	} else {
		hostCls, err := c.clusterForHost(hostname)
		var notFound *NotFoundError
		if errors.As(err, &notFound) {
			hostCls = c.makeClusterForHost(hostname)
		} else if err != nil {
			return err
		}
		return hostCls.SetStringValue(key, value)
	}
}

func (c *fileCluster) UnsetHost(hostname string) {
	if hostname == "" {
		return
	}

	hostsEntry, err := c.FindEntry("hosts")
	if err != nil {
		return
	}

	cm := ClusterMap{hostsEntry.ValueNode}
	cm.RemoveEntry(hostname)
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

func (c *fileCluster) CheckWriteable(hostname, key string) error {
	// TODO: check filesystem permissions
	return nil
}

func (c *fileCluster) Write() error {
	mainData := yaml.Node{Kind: yaml.MappingNode}
	hostsData := yaml.Node{Kind: yaml.MappingNode}

	nodes := c.documentRoot.Content[0].Content
	for i := 0; i < len(nodes)-1; i += 2 {
		if nodes[i].Value == "hosts" {
			hostsData.Content = append(hostsData.Content, nodes[i+1].Content...)
		} else {
			mainData.Content = append(mainData.Content, nodes[i], nodes[i+1])
		}
	}

	mainBytes, err := yaml.Marshal(&mainData)
	if err != nil {
		return err
	}

	filename := ClusterFile()
	err = WriteClusterFile(filename, yamlNormalize(mainBytes))
	if err != nil {
		return err
	}

	hostsBytes, err := yaml.Marshal(&hostsData)
	if err != nil {
		return err
	}

	return WriteClusterFile(HostsClusterFile(), yamlNormalize(hostsBytes))
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

// Hosts returns a list of all known hostnames clusterured in hosts.yml
func (c *fileCluster) Hosts() ([]string, error) {
	entries, err := c.hostEntries()
	if err != nil {
		return nil, err
	}

	hostnames := []string{}
	for _, entry := range entries {
		hostnames = append(hostnames, entry.Host)
	}

	sort.SliceStable(hostnames, func(i, j int) bool { return hostnames[i] == ghinstance.Default() })

	return hostnames, nil
}

func (c *fileCluster) DefaultHost() (string, error) {
	val, _, err := c.DefaultHostWithSource()
	return val, err
}

func (c *fileCluster) DefaultHostWithSource() (string, string, error) {
	hosts, err := c.Hosts()
	if err == nil && len(hosts) == 1 {
		return hosts[0], HostsClusterFile(), nil
	}

	return ghinstance.Default(), "", nil
}

func (c *fileCluster) makeClusterForHost(hostname string) *HostCluster {
	hostRoot := &yaml.Node{Kind: yaml.MappingNode}
	hostCls := &HostCluster{
		Host:      hostname,
		ClusterMap: ClusterMap{Root: hostRoot},
	}

	var notFound *NotFoundError
	hostsEntry, err := c.FindEntry("hosts")
	if errors.As(err, &notFound) {
		hostsEntry.KeyNode = &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: "hosts",
		}
		hostsEntry.ValueNode = &yaml.Node{Kind: yaml.MappingNode}
		root := c.Root()
		root.Content = append(root.Content, hostsEntry.KeyNode, hostsEntry.ValueNode)
	} else if err != nil {
		panic(err)
	}

	hostsEntry.ValueNode.Content = append(hostsEntry.ValueNode.Content,
		&yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: hostname,
		}, hostRoot)

	return hostCls
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
		return nil, errors.New("could not find any host clusterurations")
	}

	return hostClusters, nil
}

func yamlNormalize(b []byte) []byte {
	if bytes.Equal(b, []byte("{}\n")) {
		return []byte{}
	}

	return b
}

func defaultFor(key string) string {
	for _, co := range clusterOptions {
		if co.Key == key {
			return co.DefaultValue
		}
	}
	return ""
}

