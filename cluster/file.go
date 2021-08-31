package cluster

import (
	yaml "gopkg.in/yaml.v3"
)

type fileCluster struct {
	ClusterMap
	documentRoot *yaml.Node
}
