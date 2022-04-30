package app

import (
	"testing"

	"github.com/cnlisea/ant/logs"
)

func TestApp_DBMySqlRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	if err = app.DBMySqlRegister("", "yueyou", "yueyou888", "", 3306, "yueyou", "utf8mb4", "10s", false, "Local", 10, 10, 600); err != nil {
		t.Fatal("mysql register fail", err)
	}

	mysqlProxy := app.ProxyDBMySQL()
	db := mysqlProxy.GetDB("")
	if db == nil {
		t.Fatal("db not found")
	}
	var count int
	if err = db.QueryRow("SELECT count(1) FROM game").Scan(&count); err != nil {
		t.Fatal("db query row scan fail", err)
	}
	t.Log("count:", count)
}
