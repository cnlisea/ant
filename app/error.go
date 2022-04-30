package app

import "errors"

var (
	ErrDiscoveryUnavailable        = errors.New("discovery: the client unavailable")
	ErrMysqlInstanceAlreadyExisted = errors.New("mysql: the instance already existed")
	ErrMongoInstanceAlreadyExisted = errors.New("mongo: the instance already existed")
	ErrRedisInstanceAlreadyExisted = errors.New("redis: the instance already existed")
)
