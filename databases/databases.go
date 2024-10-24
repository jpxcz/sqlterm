package databases

import (
	"errors"
)

var databases = make(map[string]*Database)

// TODO: remove hardcoded values. We should get them from somewhere a file or something
func NewDatabases() map[string]*Database {
	databases = map[string]*Database{
		"Mysql-Db1": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName:    "MySql-Db1",
				Username:     "user",
				DatabaseName: "mysql_db1",
				Hostname:     "0.0.0.0",
				Password:     "password",
				Port:         "3306",
				Type:         "mysql",
			},
		},
		"Mysql-Db2": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName:    "Mysql-Db2",
				Username:     "user",
				DatabaseName: "mysql_db2",
				Hostname:     "0.0.0.0",
				Password:     "password",
				Port:         "3307",
				Type:         "mysql",
			},
		},
		"Pg-Db1": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName:    "Pg-Db1",
				Username:     "user",
				DatabaseName: "postgres_db1",
				Hostname:     "0.0.0.0",
				Password:     "password",
				Port:         "5431",
				Type:         "postgres",
			},
		},
		"SQLite-1": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName:    "SQLite-1",
				Username:     "",
				DatabaseName: "",
				Hostname:     "/Users/jp/code/sqlterm/initial_db/sqlite/db1.db",
				Password:     "",
				Port:         "",
				Type:         "sqlite3",
			},
		},
	}

	return databases
}

func GetDatabaseCredentials(key string) (DatabaseCredentials, error) {
	if databases[key] == nil {
		return DatabaseCredentials{}, errors.New("could not find database credentials for " + key)
	}

	return databases[key].DatabaseCredentials, nil
}

func GetDatabases() map[string]*Database {
	return databases
}

func GetDatabase(key string) (*Database, error) {
	if databases[key] == nil {
		return nil, errors.New("could not find database for " + key)
	}

	return databases[key], nil
}
