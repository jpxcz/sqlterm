package databases

import (
	"database/sql"
	"errors"

	"github.com/jpxcz/sqlterm/databases/mysql"
)

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

	return nil, errors.New("Database type " + dbType + " not supported")
}

func ConnectToDatabase(
	key string,
	username string,
	host string,
	port string,
	password string,
	dbType string,
) (DbConnectionStatus, error) {
    if databases[key] == nil {
       return DbDisconnected, errors.New("database could not be found to connect. Db name:"+key)
    }

	if databases[key].Db != nil && databases[key].ConnectionStatus == DbConnected {
		return databases[key].ConnectionStatus, nil
	}

	databases[key].ConnectionStatus = DbDisconnected

	databaseCredentials := databases[key].DatabaseCredentials

	db, err := createDBConnection(
		databaseCredentials.Username,
		databaseCredentials.Hostname,
		databaseCredentials.Port,
		databaseCredentials.Password,
		databaseCredentials.Type,
	)
	if err != nil {
		databases[key].ConnectionStatus = DbErrorConnection
	} else {
		databases[key].Db = db
		databases[key].ConnectionStatus = DbConnected
	}

	return databases[key].ConnectionStatus, err
}

func DisconnectFromDatabase(key string) error {
    if databases[key] == nil {
		return errors.New("database could not be found to attempt a disconnect. Db name:"+key)
    }

    if databases[key].Db == nil {
        return nil
    }

    err := databases[key].Db.Close()
    databases[key].Db = nil
    if err != nil {
        databases[key].ConnectionStatus = DbErrorConnection
        return err
    }

    databases[key].ConnectionStatus = DbDisconnected
    return nil
}

func GetConnection(key string) *Database {
	if databases[key] == nil {
		return nil
	}

	return databases[key]
}
