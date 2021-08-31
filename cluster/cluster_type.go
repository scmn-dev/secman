package cluster

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
