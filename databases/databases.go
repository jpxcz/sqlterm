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

type Database struct {
	Db                  *sql.DB
	ConnectionStatus    DbConnectionStatus
	DatabaseCredentials DatabaseCredentials
}

type DatabaseCredentials struct {
	ShortName string `json:"shortname"`
	Username  string `json:"username"`
	Hostname  string `json:"hostname"`
	Password  string `json:"password"`
	Port      string `json:"port"`
	Type      string `json:"type"`
}

var databases = make(map[string]*Database)

func createDBConnection(
	username string,
	host string,
	port string,
	password string,
	dbType string,
) (*sql.DB, error) {
	if dbType == "mysql" {
		return mysql.CreateDBConnection(username, host, port, password)
	}

	return nil, errors.New("Database type not supported")
}

func ConnectToDatabase(
	key string,
	username string,
	host string,
	port string,
	password string,
	dbType string,
) (*Database, error) {
	if databases[key] != nil && databases[key].Db != nil {
		return databases[key], nil
	}

	connection := &Database{
		Db:               nil,
		ConnectionStatus: DbDisconnected,
	}

	databases[key] = connection
	db, err := createDBConnection(username, host, port, password, dbType)
	if err != nil {
		connection.ConnectionStatus = DbErrorConnection
	} else {
		connection.Db = db
		connection.ConnectionStatus = DbConnected
	}

	databases[key] = connection
	return connection, err
}

func DisconnectFromDatabase(key string) {
	if databases[key] == nil {
		return
	}

	databases[key].Db.Close()
	delete(databases, key)
}

func GetConnection(key string) *Database {
	if databases[key] == nil {
		return nil
	}

	return databases[key]
}
