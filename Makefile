
build:
	go get github.com/mattn/go-sqlite3
	go build --tags "icu json1 fts5 secure_delete userauth" github.com/mattn/go-sqlite3


install:
	go install github.com/mattn/go-sqlite3


.PHONY: build install
