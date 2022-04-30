package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Client struct {
	client *redis.Pool
}

func NewClient(password string, addr string, port uint16, db int,
	active int, idle int, idleTimeout int) (*Client, error) {
	client := &redis.Pool{
		MaxIdle:     idle,
		MaxActive:   active,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr+":"+strconv.Itoa(int(port)))
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err = c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			_, err = c.Do("SELECT", db)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
	return &Client{client: client}, nil
}

func (c *Client) GetConn(ctx context.Context) redis.Conn {
	if c.client == nil {
		return nil
	}
	return c.client.Get()
}

func (c *Client) Close(ctx context.Context) error {
	if c.client == nil {
		return nil
	}
	return c.client.Close()
}
