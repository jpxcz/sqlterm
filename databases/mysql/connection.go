package mysql

import (
	"database/sql"

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
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

    // do not close it. It must be handled after the queries are finished
	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
