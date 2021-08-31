package cluster

type Config interface {
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
