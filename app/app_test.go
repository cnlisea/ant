package app

import (
	"testing"
)

func TestNew(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Run(); err != nil {
		t.Fatal(err)
	}
}
