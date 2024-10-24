package databases

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jpxcz/sqlterm/databases/mysql"
	"github.com/jpxcz/sqlterm/databases/postgress"
	sqlite3 "github.com/jpxcz/sqlterm/databases/sqlite"
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
	} else if d.DatabaseCredentials.Type == "postgres" {
		return postgress.CreateDBConnection(
			d.DatabaseCredentials.Username,
			d.DatabaseCredentials.DatabaseName,
			d.DatabaseCredentials.Hostname,
			d.DatabaseCredentials.Port,
			d.DatabaseCredentials.Password,
		)
	} else if d.DatabaseCredentials.Type == "sqlite3" {
		return sqlite3.CreateDBConnection(
			d.DatabaseCredentials.Hostname,
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

	log.Println("running query", query)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
