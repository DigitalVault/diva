package db

import (
  "testing"
  "os"
)

func TestInit(t *testing.T) {
  var db = DbSqlite { "./test.sqlite", nil }
  if db.init() != nil {
    t.Error("Expected error return value to be nil.")
  }
  os.Remove("./test.sqlite")
}
