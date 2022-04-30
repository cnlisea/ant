package ws

import (
	"context"
	"net/http"

	"github.com/cnlisea/ant/logs"
	"github.com/gorilla/websocket"
)

func Handler(ctx context.Context, origin bool, heartbeatInterval int64, receive func(*Connection)) func(w http.ResponseWriter, r *http.Request) {
	upgrade := &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}
	if origin {
		upgrade.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			logs.Warn("up grade fail", logs.Error("err", err))
			return
		}
		receive(NewConnection(ctx, wsConn, ClientIp(r), heartbeatInterval))
	}
}
