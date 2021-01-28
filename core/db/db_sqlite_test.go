package db

import (
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetReportCaller(true)
	os.Exit(m.Run())
}

func TestInit(t *testing.T) {
	os.Remove("./test.sqlite")
	var db = DbSqlite{Path : "./test.sqlite"}
	if db.init() != nil {
		t.Error("Expected error return value to be nil.")
	}
	os.Remove("./test.sqlite")
}

