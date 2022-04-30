package mongo

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	dbName string
	client *mongo.Client
}

func NewClient(ctx context.Context,
	user string, password string,
	addr []string, dbName string,
	replicaSet string, connTimeout int,
	active int, idle int, idleTimeout int) (*Client, error) {
	var (
		dsn      strings.Builder
		userAuth bool
	)
	dsn.WriteString("mongodb://")
	if user != "" && password != "" {
		dsn.WriteString(user)
		dsn.WriteString(":")
		dsn.WriteString(password)
		dsn.WriteString("@")
		userAuth = true
	}
	dsn.WriteString(strings.Join(addr, ","))
	dsn.WriteString("/")
	dsn.WriteString("?w=majority&authSource=")
	dsn.WriteString(dbName)
	if userAuth {
		dsn.WriteString("&authMechanism=SCRAM-SHA-1")
	}
	if connTimeout > 0 {
		dsn.WriteString("&connectTimeoutMS=")
		dsn.WriteString(strconv.Itoa(connTimeout * 1000))
	}
	if active > 0 {
		dsn.WriteString("&maxPoolSize=")
		dsn.WriteString(strconv.Itoa(active))
	}
	if idle > 0 {
		dsn.WriteString("&minPoolSize=")
		dsn.WriteString(strconv.Itoa(idle))
	}
	if idleTimeout > 0 {
		dsn.WriteString("&maxIdleTimeMS=")
		dsn.WriteString(strconv.Itoa(idleTimeout * 1000))
	}
	if replicaSet != "" {
		dsn.WriteString("&replicaSet=")
		dsn.WriteString(replicaSet)
	}

	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn.String()))
	if err != nil {
		return nil, errors.New("client connect fail, err:" + err.Error())
	}

	ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		client.Disconnect(ctx)
		return nil, errors.New("client ping fail, err:" + err.Error())
	}

	return &Client{
		dbName: dbName,
		client: client,
	}, nil
}

func (c *Client) GetConn(ctx context.Context) *mongo.Database {
	if c.client == nil {
		return nil
	}

	return c.client.Database(c.dbName)
}

func (c *Client) Close(ctx context.Context) error {
	if c.client == nil {
		return nil
	}

	if ctx == nil {
		ctx = context.Background()
	}
	return c.client.Disconnect(ctx)
}
