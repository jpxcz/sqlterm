package postgress

import (
	"database/sql"
	"log"

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
    log.Println("Connecting to Postgres database with DSN", dns)
	db, err := sql.Open("postgres", dns)
	if err != nil {
        log.Println("Error opening Postgres database", err)
		return nil, err
	}

	// do not close it. It must be handled after the queries are finished
	// defer db.Close()

    log.Println("Pinging Postgres database")

	if err := db.Ping(); err != nil {
        log.Println("Error pinging Postgres database", err)
		return nil, err
	}
    log.Println("Connected to Postgres database")

	return db, nil
}
