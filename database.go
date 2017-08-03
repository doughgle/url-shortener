package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Database interface
type Database interface {
	Get(shortened string) (string, error)
	Save(shortened string, url string, user_id int) (string, error)
}

type sqlite struct {
	Path string
}

func (s sqlite) Save(shortened string, url string, user_id int) (string, error) {
	db, err := sql.Open("sqlite3", s.Path)
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := tx.Prepare("insert into urls(shortened, url, user_id) values(?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	result, err := stmt.Exec(shortened, url, user_id)
	if err != nil {
		return "", err
	}
	_ = result

	tx.Commit()

	return shortened, nil
}

func (s sqlite) Get(shortened string) (string, error) {
	db, err := sql.Open("sqlite3", s.Path)
	stmt, err := db.Prepare("select url from urls where shortened = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var url string
	err = stmt.QueryRow(shortened).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s sqlite) Init() {
	c, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	sqlStmt := `create table if not exists urls (shortened text not null primary key, url text not null, user_id integer, created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL);`
	_, err = c.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
