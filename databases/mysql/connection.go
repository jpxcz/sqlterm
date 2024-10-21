package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDBConnection(
	username string,
	databaseName string,
	host string,
	port string,
	password string,
) (*sql.DB, error) {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName
	log.Println("connecting to database", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
