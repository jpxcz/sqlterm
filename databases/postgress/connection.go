package postgress

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateDBConnection(
	username string,
	databaseName string,
	host string,
	port string,
	password string,
) (*sql.DB, error) {
	dns := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + databaseName + "?sslmode=disable"
	db, err := sql.Open("postgres", dns)
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
