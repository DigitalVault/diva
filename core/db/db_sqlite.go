package db

import (
    "database/sql"
//    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
//    "os"
)

type DbSqlite struct {
  Path string
  Db *sql.DB 
}

func (dbSqlite *DbSqlite)init() (error) {
  log.Printf("[DbSqlite::init] Initialising digital vault DB at %s", dbSqlite.Path)

	db, err := sql.Open("sqlite3", dbSqlite.Path)
  dbSqlite.Db = db
	if err != nil {
		log.Fatal(err)
    return err
	}
	defer dbSqlite.Db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = dbSqlite.Db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

  return nil;
}

