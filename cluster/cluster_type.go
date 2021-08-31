package cluster

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"
)

type Cluster interface {
	Get(string, string) (string, error)
	GetWithSource(string, string) (string, string, error)
	Set(string, string, string) error
	UnsetHost(string)
	Hosts() ([]string, error)
	DefaultHost() (string, error)
	DefaultHostWithSource() (string, string, error)
	CheckWriteable(string, string) error
	Write() error
}

type ClusterOption struct {
	Key           string
	Description   string
	DefaultValue  string
	AllowedValues []string
}

var clusterOptions = []ClusterOption{
	{
		Key:           "git_protocol",
		Description:   "the protocol to use for git clone and push operations",
		DefaultValue:  "https",
		AllowedValues: []string{"https", "ssh"},
	},
	{
		Key:           "prompt",
		Description:   "toggle interactive prompting in the terminal",
		DefaultValue:  "enabled",
		AllowedValues: []string{"enabled", "disabled"},
	},
	{
		Key:          "http_unix_socket",
		Description:  "the path to a unix socket through which to make HTTP connection",
		DefaultValue: "",
	},
	{
		Key:          "browser",
		Description:  "the web browser to use for opening URLs",
		DefaultValue: "",
	},
}

func ClusterOptions() []ClusterOption {
	return clusterOptions
}

func ValidateKey(key string) error {
	for _, configKey := range clusterOptions {
		if key == configKey.Key {
			return nil
		}
	}

	return fmt.Errorf("invalid key")
}

type InvalidValueError struct {
	ValidValues []string
}

func (e InvalidValueError) Error() string {
	return "invalid value"
}

func ValidateValue(key, value string) error {
	var validValues []string

	for _, v := range clusterOptions {
		if v.Key == key {
			validValues = v.AllowedValues
			break
		}
	}

	if validValues == nil {
		return nil
	}

	for _, v := range validValues {
		if v == value {
			return nil
		}
	}

	return &InvalidValueError{ValidValues: validValues}
}

func NewCluster(root *yaml.Node) Cluster {
	return &fileCluster{
		ClusterMap:    ClusterMap{Root: root.Content[0]},
		documentRoot: root,
	}
}

// NewFromString initializes a Cluster from a yaml string
func NewFromString(str string) Cluster {
	root, err := parseClusterData([]byte(str))
	if err != nil {
		panic(err)
	}

	return NewCluster(root)
}

// NewBlankCluster initializes a config file pre-populated with comments and default values
func NewBlankCluster() Cluster {
	return NewCluster(NewBlankRoot())
}

func NewBlankRoot() *yaml.Node {
	return &yaml.Node{
		Kind: yaml.DocumentNode,
		Content: []*yaml.Node{
			{
				Kind: yaml.MappingNode,
				Content: []*yaml.Node{
					{
						HeadComment: "What protocol to use when performing git operations. Supported values: ssh, https",
						Kind:        yaml.ScalarNode,
						Value:       "git_protocol",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "https",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "",
					},
					{
						HeadComment: "When to interactively prompt. This is a global config that cannot be overridden by hostname. Supported values: enabled, disabled",
						Kind:        yaml.ScalarNode,
						Value:       "prompt",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "enabled",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "",
					},
					{
						HeadComment: "The path to a unix socket through which send HTTP connections. If blank, HTTP traffic will be handled by net/http.DefaultTransport.",
						Kind:        yaml.ScalarNode,
						Value:       "http_unix_socket",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "",
					},
					{
						HeadComment: "What web browser secman should use when opening URLs. If blank, will refer to environment.",
						Kind:        yaml.ScalarNode,
						Value:       "browser",
					},
				},
			},
		},
	}
}
