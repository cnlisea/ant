package rpc

type ClientOption struct {
	SelectMode       ClientSelectMode
	GroupName        string
	ServerLocalProxy ServerLocalProxy
}

type ClientOptionFunc func(*ClientOption)

type ClientSelectMode uint8

const (
	ClientSelectModeRoundRobin ClientSelectMode = iota
	ClientSelectModeHash
	ClientSelectModeSoftState
)

func ClientWithSelectMode(selectMode ClientSelectMode) ClientOptionFunc {
	return func(op *ClientOption) {
		op.SelectMode = selectMode
	}
}

func ClientWithGroupName(groupName string) ClientOptionFunc {
	return func(op *ClientOption) {
		op.GroupName = groupName
	}
}

func ClientWithServerLocalProxy(serverLocalProxy ServerLocalProxy) ClientOptionFunc {
	return func(op *ClientOption) {
		op.ServerLocalProxy = serverLocalProxy
	}
}
