package ws

import "time"

func (c *Connection) HeartbeatChecker() {
	var (
		timer = time.NewTimer(time.Duration(c.heartbeatInterval) * time.Second)
	)
	for {
		select {
		case <-timer.C:
			if !c.IsAlive() {
				c.Close()
				goto EXIT
			}
			timer.Reset(time.Duration(c.heartbeatInterval) * time.Second)
		case <-c.ctx.Done():
			timer.Stop()
			goto EXIT
		}
	}

EXIT:
}

func (c *Connection) IsAlive() bool {
	alive := true

	c.mutex.RLock()
	if c.isClosed || time.Now().Unix()-c.lastHeartbeatTime > c.heartbeatInterval {
		alive = false
	}
	c.mutex.RUnlock()

	return alive
}

func (c *Connection) KeepAlive() {
	c.mutex.Lock()
	c.lastHeartbeatTime = time.Now().Unix()
	c.mutex.Unlock()
}
