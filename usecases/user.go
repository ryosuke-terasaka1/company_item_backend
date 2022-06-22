package usecases

import (
	"database/sql"
	"log"
)

type DB struct {
	name string
	url  string
}

var db DB

var DbConnection, _ = sql.Open(db.name, db.url)

func create_table() {
	cmd := `CREATE TABLE IF NOT EXISTS user(
    name STRING,
    email STRING)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
