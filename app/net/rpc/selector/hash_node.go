package selector

import "strings"

type HashNode struct {
	Server string
	Tag    string
}

func createHashNode(servers map[string]string) []*HashNode {
	var (
		ss   = make([]*HashNode, 0, len(servers))
		key  string
		keys []string
		node *HashNode
	)
	for key = range servers {
		node = &HashNode{Server: key}
		keys = strings.SplitN(key, "@", 2)
		if len(keys) > 1 {
			node.Tag = keys[1]
		}
		ss = append(ss, node)
	}

	return ss
}
