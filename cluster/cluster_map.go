package cluster

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type ClusterMap struct {
	Root *yaml.Node
}

type ClusterEntry struct {
	KeyNode   *yaml.Node
	ValueNode *yaml.Node
	Index     int
}

type NotFoundError struct {
	error
}

func (cm *ClusterMap) Empty() bool {
	return cm.Root == nil || len(cm.Root.Content) == 0
}

func (cm *ClusterMap) GetStringValue(key string) (string, error) {
	entry, err := cm.FindEntry(key)
	if err != nil {
		return "", err
	}

	return entry.ValueNode.Value, nil
}

func (cm *ClusterMap) SetStringValue(key, value string) error {
	entry, err := cm.FindEntry(key)

	var notFound *NotFoundError

	valueNode := entry.ValueNode

	if err != nil && errors.As(err, &notFound) {
		keyNode := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: key,
		}

		valueNode = &yaml.Node{
			Kind:  yaml.ScalarNode,
			Tag:   "!!str",
			Value: "",
		}

		cm.Root.Content = append(cm.Root.Content, keyNode, valueNode)
	} else if err != nil {
		return err
	}

	valueNode.Value = value

	return nil
}

func (cm *ClusterMap) FindEntry(key string) (ce *ClusterEntry, err error) {
	err = nil

	ce = &ClusterEntry{}

	// Content slice goes [key1, value1, key2, value2, ...]
	topLevelPairs := cm.Root.Content
	for i, v := range topLevelPairs {
		// Skip every other slice item since we only want to check against keys
		if i%2 != 0 {
			continue
		}
		if v.Value == key {
			ce.KeyNode = v
			ce.Index = i
			if i+1 < len(topLevelPairs) {
				ce.ValueNode = topLevelPairs[i+1]
			}

			return
		}
	}

	return ce, &NotFoundError{errors.New("not found")}
}

func (cm *ClusterMap) RemoveEntry(key string) {
	newContent := []*yaml.Node{}

	content := cm.Root.Content
	for i := 0; i < len(content); i++ {
		if content[i].Value == key {
			i++ // skip the next node which is this key's value
		} else {
			newContent = append(newContent, content[i])
		}
	}

	cm.Root.Content = newContent
}
