package db

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	var db = DbSqlite{Path : "./test.sqlite"}
	if db.init() != nil {
		t.Error("Expected error return value to be nil.")
	}
	os.Remove("./test.sqlite")
}
