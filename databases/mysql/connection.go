package mysql

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func CreateDBConnection(username string, host string, port string, password string) (*sql.DB, error) {
    dsn := username+":"+password+"@tcp("+host+":"+port+")/"
    log.Println("connecting to database", dsn)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}


