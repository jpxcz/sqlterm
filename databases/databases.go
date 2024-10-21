package databases

import (
	"errors"
)

var databases = make(map[string]*Database)

// TODO: remove hardcoded values. We should get them from somewhere a file or something
func NewDatabases() map[string]*Database {
	databases = map[string]*Database{
		"DB1": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName: "DB1",
				Username:  "user1",
                DatabaseName: "db1",
				Hostname:  "0.0.0.0",
				Password:  "password1",
				Port:      "3306",
				Type:      "mysql",
			},
		},
		"DB2": {
			ConnectionStatus: DbDisconnected,
			DatabaseCredentials: DatabaseCredentials{
				ShortName: "DB2",
				Username:  "user2",
                DatabaseName: "db2",
				Hostname:  "0.0.0.0",
				Password:  "password2",
				Port:      "3307",
				Type:      "mysql",
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
