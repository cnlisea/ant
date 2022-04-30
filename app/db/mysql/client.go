package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	client *sql.DB
}

func NewClient(user string, password string,
	addr string, port uint16, dbName string,
	charset string, connTimeout string, parseTime bool, loc string,
	active int, idle int, idleTimeout int) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=%t&loc=%s",
		user, password, addr, port, dbName, charset, connTimeout, parseTime, loc)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("client open fail " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, errors.New("client ping fail " + err.Error())
	}

	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(active)
	db.SetConnMaxLifetime(time.Duration(idleTimeout) * time.Second)
	return &Client{client: db}, nil
}

func (c *Client) GetConn(ctx context.Context) *sql.DB {
	return c.client
}

func (c *Client) Close(ctx context.Context) error {
	if c.client == nil {
		return nil
	}
	return c.client.Close()
}
