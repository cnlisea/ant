package ws

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn              *websocket.Conn
	realIP            string
	inChan            chan []byte
	outChan           chan []byte
	ctx               context.Context
	cancel            context.CancelFunc
	mutex             *sync.RWMutex
	isClosed          bool
	heartbeatInterval int64
	lastHeartbeatTime int64
	err               error
}

func NewConnection(ctx context.Context, conn *websocket.Conn, realIP string, heartbeatInterval int64) *Connection {
	if heartbeatInterval <= 0 {
		// default not check
		heartbeatInterval = 0
	}
	c := &Connection{
		conn:              conn,
		realIP:            realIP,
		inChan:            make(chan []byte, 100),
		outChan:           make(chan []byte, 100),
		mutex:             new(sync.RWMutex),
		heartbeatInterval: heartbeatInterval,
		lastHeartbeatTime: time.Now().Unix(),
	}
	if ctx == nil {
		ctx = context.Background()
	}
	c.ctx, c.cancel = context.WithCancel(ctx)

	go c.readLoop()
	go c.writeLoop()

	if c.heartbeatInterval != 0 {
		go c.HeartbeatChecker()
	}

	return c
}

func (c *Connection) ReadMessage() ([]byte, error) {
	var (
		err  error
		data []byte
	)
	select {
	case data = <-c.inChan:
	case <-c.ctx.Done():
		err = ErrConnectionClosed
		if c.err != nil {
			err = c.err
		}
	}

	return data, err
}

func (c *Connection) WriteMessage(data []byte) error {
	var err error

	select {
	case c.outChan <- data:
	case <-c.ctx.Done():
		err = ErrConnectionClosed
	}

	return err
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) RemoteRealIP() string {
	return c.realIP
}

func (c *Connection) Close() {
	c.mutex.Lock()
	if !c.isClosed {
		c.cancel()
		c.isClosed = true
	}
	c.mutex.Unlock()
}
