package selector

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/cnlisea/ant/logs"
	"github.com/valyala/fastrand"
)

type SoftStateNode struct {
	Server       string
	Games        []string
	AbandonArray [10]byte
}

func createSoftStateNode(servers map[string]string) []*SoftStateNode {
	var (
		ss     = make([]*SoftStateNode, 0, len(servers))
		node   *SoftStateNode
		k, v   string
		values url.Values
		value  string
		weight int
		games  []string
		err    error
	)
	for k, v = range servers {
		games = nil
		weight = 10
		if values, err = url.ParseQuery(v); err == nil {
			if value = values.Get("offline"); value != "" &&
				value != "0" &&
				value != "false" {
				continue
			}

			if value = values.Get("games"); value != "" {
				games = strings.Split(value, ",")
			}

			if value = values.Get("weight"); value != "" {
				if weight, err = strconv.Atoi(value); err != nil {
					logs.Warn("server mate data weight invalid", logs.String("weight", value), logs.String("server", k))
					weight = 10
				}
			}
		}
		node = &SoftStateNode{
			Server:       k,
			Games:        games,
			AbandonArray: randomGenerateGameNodeAbandonArray(weight),
		}
		ss = append(ss, node)
	}

	return ss
}

func randomGenerateGameNodeAbandonArray(weight int) [10]byte {
	var ret [10]byte
	if weight >= 10 {
		return ret
	}

	var i int
	if weight <= 0 {
		for i = 0; i < 10; i++ {
			ret[i] = 1
		}
		return ret
	}

	var index uint32
	for i = 0; i < 10-weight; i++ {
		index = fastrand.Uint32() % 10
		for ret[index] == 1 {
			index = fastrand.Uint32() % 10
		}
		ret[index] = 1
	}

	return ret
}
