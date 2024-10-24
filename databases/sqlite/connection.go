package sqlite3

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDBConnection(
	filePath string,
) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
        log.Println("Error opening sqlite3 database", err)
		return nil, err
	}

	// do not close it. It must be handled after the queries are finished
	// defer db.Close()

	if err := db.Ping(); err != nil {
	       log.Println("Error pinging sqlite3 database", err)
		return nil, err
	}

	return db, nil
}
