package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ConfigDB struct {
	Driver string
	Path   string
	Name   string
}

func NewConfDB() *ConfigDB {
	return &ConfigDB{
		Driver: "sqlite3",
		Name:   "forum.db",
	}
}

func InitDB(c *ConfigDB) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, c.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	file, err := ioutil.ReadFile("./migrations/db.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(file)); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
