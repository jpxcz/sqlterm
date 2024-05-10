package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jpxcz/sqlterm/databases"
	mysqlclient "github.com/jpxcz/sqlterm/mysql_client"
)

func main() {
	// Get the database info
	dbs, databaseMap, databaseKeys := databases.GetDatabases()

	// Construct the command line args
	var dbEnv string
	flag.StringVar(&dbEnv, "env", dbs[0].Key, "One of the following database environment keys; ["+strings.Join(databaseKeys, ",")+"]")

	var dbTableFormat string
	flag.StringVar(&dbTableFormat, "table", "NO", "Format SQL tables [YES / NO]")
	flag.Parse()

	// Create a map of args that were supplied on the command line
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if flagset["env"] {
		// If the environment was specified on the command line then use it
		dbCreds := databaseMap[dbEnv]
		mysqlclient.ExecMySqlClient(dbCreds.Username, dbCreds.Hostname, dbCreds.Password, flagset["table"])
	} else {
		// If the option wasn't supplioed on the command line then ask for it
		fmt.Println("Welcome, please select one of the databases to connect")
		for i, db := range dbs {
			fmt.Printf("[%d] %s - %s\n", i, db.Key, db.ShortName)
		}

		var i int
		fmt.Scanf("%d\n", &i)
		if i >= len(dbs) {
			log.Fatalf("Exiting: option [%d] selected is not in range of the databases. Exiting application\n", i)
			os.Exit(1)
		}

		dbCreds := dbs[i]
		mysqlclient.ExecMySqlClient(dbCreds.Username, dbCreds.Hostname, dbCreds.Password, flagset["table"])
	}

}
