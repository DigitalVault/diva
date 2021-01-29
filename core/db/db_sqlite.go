package db

import (
	"database/sql"
	//    "fmt"
	_ "github.com/mattn/go-sqlite3"
	//"log"
	log "github.com/sirupsen/logrus"
	//    "os"
)

type DbSqlite struct {
	Path string
	Db   *sql.DB
}

func (dbSqlite *DbSqlite) init() error {
	log.Infof("Initialising digital vault DB at %s", dbSqlite.Path)

	db, err := sql.Open("sqlite3", dbSqlite.Path)
	dbSqlite.Db = db
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer dbSqlite.Db.Close()

	/*
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	*/

	initSqlStmt := `
	create table source_files (
		id integer primary key autoincrement,
		hostname text not null, 
		drive text not null, 
		path text not null, 
		filename text not null, 
		date_modified datetime not null, 
		filetype varchar(20), 
		work_status varchar(20), 
		unique (hostname, drive, path, filename)
	);
	`
	_, err = dbSqlite.Db.Exec(initSqlStmt)
	if err != nil {
		log.Infof("%q: %s\n", err, initSqlStmt)
		return err
	}

	return nil
}
