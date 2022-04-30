package app

import (
	"net/http"
	"testing"

	"github.com/cnlisea/ant/logs"
	"github.com/go-chi/chi"
)

func TestApp_NetHttpRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	var router = chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("app http"))
	})
	if err = app.NetHttpRegister("", "", 8080, nil, router); err != nil {
		t.Fatal(err)
	}
	if err = app.Run(); err != nil {
		t.Fatal(err)
	}
}
