package selector

import "context"

const (
	CtxSoftStateTag = "ctx_soft_state_tag"
)

type SoftState struct {
	servers     []*SoftStateNode
	serverTimes []uint64
	times       uint64
}

func (s *SoftState) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	var ss = s.servers
	if len(ss) == 0 {
		return ""
	}

	playTag, ok := ctx.Value(CtxSoftStateTag).(string)
	if !ok {
		return ""
	}

	var (
		optionList = make([]int, 0, len(s.servers))
		i, j       int
	)
	for i = range s.servers {
		if len(s.servers[i].Games) == 0 {
			optionList = append(optionList, i)
			continue
		}

		for j = range s.servers[i].Games {
			if s.servers[i].Games[j] == playTag {
				optionList = append(optionList, i)
				break
			}
		}
	}

	count := uint64(len(optionList))
	if count == 0 {
		return ""
	}

	var (
		useServerIndex = -1
		index          uint64
	)
	for i := s.times; i < s.times+count; i++ {
		index = i % count
		if s.servers[optionList[index]].AbandonArray[s.serverTimes[optionList[index]]%10] == 0 {
			useServerIndex = optionList[index]
			break
		}
		s.serverTimes[optionList[index]]++
	}

	// 未选出, 按最小调用次数获取
	if useServerIndex == -1 {
		useServerIndex = optionList[0]
		var minTime = s.serverTimes[useServerIndex]
		for i = 1; i < len(optionList); i++ {
			if s.serverTimes[optionList[i]] < minTime {
				useServerIndex = optionList[i]
				minTime = s.serverTimes[useServerIndex]
			}
		}
	}

	s.serverTimes[useServerIndex]++
	s.times++
	return s.servers[useServerIndex].Server
}

func (s *SoftState) UpdateServer(servers map[string]string) {
	ss := createSoftStateNode(servers)
	s.servers = ss
	s.serverTimes = make([]uint64, len(ss))
	s.times = 0
}
