package proxy

import "go.mongodb.org/mongo-driver/mongo"

type DBMongo interface {
	GetDB(name ...string) *mongo.Database
}
