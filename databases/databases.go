package databases

import (
	"fmt"
	"os"
)

type DatabaseCredentials struct {
	Key       string `json:"key"`
	ShortName string `json:"shortname"`
	Username  string `json:"username"`
	Hostname  string `json:"hostname"`
	Password  string `json:"password"`
	Port      string `json:"port"`
}

func GetDatabases() ([]DatabaseCredentials, map[string]DatabaseCredentials, []string) {
	// read the config file containing the database credentials
	databases, err := ReadDatabasesJson()
	if err != nil {
		fmt.Printf("Exiting: could not read databases file correctly. %v\n", err)
		os.Exit(1)
	} else if len(databases) == 0 {
		fmt.Println("Exiting: database config file [~/.config/sqlterm/databases.json] contained no data")
		os.Exit(1)
	}

	// Create a map of db envs to login credentials
	databaseMap := make(map[string]DatabaseCredentials)

	// Create an array of the db env keys (used for the cmdline args help)
	var databaseKeys []string

	for _, db := range databases {
		// fillout the db envs map and the array
		databaseMap[db.Key] = db
		databaseKeys = append(databaseKeys, db.Key)
	}

	return databases, databaseMap, databaseKeys
}
