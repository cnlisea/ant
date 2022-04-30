package selector

import "context"

const CtxHashTag = "ctx_hash_tag"

type Hash struct {
	servers []*HashNode
}

func (h *Hash) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	var ss = h.servers
	if len(ss) == 0 {
		return ""
	}

	value, ok := ctx.Value(CtxHashTag).(string)
	if !ok {
		return ""
	}

	for i := range ss {
		if ss[i] != nil && ss[i].Tag == value {
			return ss[i].Server
		}
	}
	return ""
}

func (h *Hash) UpdateServer(servers map[string]string) {
	ss := createHashNode(servers)
	h.servers = ss
}
