package main

import (
	// "fmt"
	"log"
	// "os"

	// "github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/tui"
	// mysqlclient "github.com/jpxcz/sqlterm/mysql_client"
)

// func main() {
// 	dbs := databases.GetDatabases()
//     fmt.Println("Welcome, please select one of the databases to connect")
// 	for i, db := range dbs {
// 		fmt.Printf("[%d] %v\n", i, db.ShortName)
// 	}
//
// 	var i int
// 	fmt.Scanf("%d\n", &i)
// 	if i >= len(dbs) {
// 		log.Fatalf("option [%d] selected is not in range of the databases. Exiting application\n", i)
// 		os.Exit(1)
// 	}
//
// 	db := dbs[i]
// 	mysqlclient.ExecMySqlClient(db.Username, db.Hostname, db.Password)
// }
//
func main() {
    p := tui.NewTeaProgram()
    if _, err := p.Run(); err != nil {
        log.Fatal(err)

    }
}
