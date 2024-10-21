package databases

import (
	"database/sql"
	"errors"

	"github.com/jpxcz/sqlterm/databases/mysql"
)

const (
	DbDisconnected DbConnectionStatus = iota
	DbConnected
	DbErrorConnection
)

type DbConnectionStatus uint

type DatabaseCredentials struct {
	ShortName    string `json:"shortname"`
	Username     string `json:"username"`
	DatabaseName string `json:"database_name"`
	Hostname     string `json:"hostname"`
	Password     string `json:"password"`
	Port         string `json:"port"`
	Type         string `json:"type"`
}

type Database struct {
	ConnectionStatus    DbConnectionStatus
	DatabaseCredentials DatabaseCredentials
}

func (d *Database) Connect() (*sql.DB, error) {
	if d.DatabaseCredentials.Type == "mysql" {
		return mysql.CreateDBConnection(
			d.DatabaseCredentials.Username,
			d.DatabaseCredentials.DatabaseName,
			d.DatabaseCredentials.Hostname,
			d.DatabaseCredentials.Port,
			d.DatabaseCredentials.Password,
		)
	}

	return nil, errors.New("database type " + d.DatabaseCredentials.Type + " not supported yet")
}

func (d *Database) Ping() error {
    db, err := d.Connect()
    defer db.Close()

    if err != nil {
        return err
    }

    err = db.Ping()
    if err != nil {
        d.ConnectionStatus = DbErrorConnection
    } else {
        d.ConnectionStatus = DbConnected
    }

    return err
}

func (d *Database) Close() {
    d.ConnectionStatus = DbDisconnected
}

func (d *Database) Query(query string) (*sql.Rows, error) {
    db, err := d.Connect()
    defer db.Close()

    if err != nil {
        return nil, err
    }

    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    return rows, nil
}
