package app

import (
	"context"
	"testing"

	"github.com/cnlisea/ant/logs"
	"go.mongodb.org/mongo-driver/bson"
)

func TestApp_DBMongoRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	if err = app.DBMongoRegister("", "yueyou", "yueyou888", []string{""}, "room", "yueyou", 10, 10, 10, 600); err != nil {
		t.Fatal("mongo register fail", err)
	}

	proxy := app.ProxyDBMongo()

	db := proxy.GetDB("")
	if db == nil {
		t.Fatal("db not found")
	}

	collection := db.Collection("room")
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		t.Fatal("count fail", err)
	}

	t.Log("count:", count)
}
