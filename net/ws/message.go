package ws

import (
	"github.com/gorilla/websocket"
	"strings"
)

func (c *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = c.conn.ReadMessage(); err != nil {
			goto ERR
		}

		select {
		case c.inChan <- data:
			if c.heartbeatInterval > 0 {
				c.KeepAlive()
			}
		case <-c.ctx.Done():
			goto ERR
		}
	}

ERR:
	if err != nil &&
		(websocket.IsCloseError(err,
			websocket.CloseNormalClosure,
			websocket.CloseGoingAway,
			websocket.CloseAbnormalClosure,
			websocket.CloseNoStatusReceived) ||
			strings.Contains(err.Error(), "connection reset by peer")) {
		err = ErrClientConnClosed
	}

	c.err = err
	c.Close()
}

func (c *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-c.outChan:
			if err = c.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				goto ERR
			}
		case <-c.ctx.Done():
			for {
				select {
				case data = <-c.outChan:
					if err = c.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
						goto ERR
					}
				default:
					goto ERR
				}
			}
		}
	}

ERR:
	c.conn.Close()
	c.Close()
}
