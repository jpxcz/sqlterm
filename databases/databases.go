package databases

import (
	"log"
	"os"
)

type DatabaseCredentials struct {
	ShortName string `json:"shortname"`
	Username  string `json:"username"`
	Hostname  string `json:"hostname"`
	Password  string `json:"password"`
	Port      string `json:"port"`
}

func GetDatabases() []DatabaseCredentials {
    databases, err := ReadDatabasesJson()
    if err != nil {
        log.Fatalf("could not read databases file correctly. %v", err )
        os.Exit(1)
    }

    return databases
}
