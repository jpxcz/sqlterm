package commands

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpxcz/sqlterm/databases"
)

type MsgDatabaseSelectionUpdate bool

func CmdDatabaseSelectionUpdate() tea.Msg {
	return MsgDatabaseSelectionUpdate(true)
}

func CmdDatabaseQuery(query string) tea.Cmd {
	return func() tea.Msg {
		dbs := databases.GetDatabases()
		for _, db := range dbs {
			if db.ConnectionStatus != databases.DbConnected {
				continue
			}
			rows, err := db.Query(query)
			if err != nil {
				log.Println(
					"Error executing query for database",
					db.DatabaseCredentials.ShortName,
					err,
				)
			}

			columns, err := rows.Columns()
			if err != nil {
				log.Println("Error getting columns", err)
			}
			columnsTypes, err := rows.ColumnTypes()

			for _, col := range columnsTypes {
				log.Println(col.Name())
			}

			values := make([]interface{}, len(columns))
			valuePointers := make([]interface{}, len(columns))
			for i := range values {
				valuePointers[i] = &values[i]
			}

			for rows.Next() {
				if err := rows.Scan(valuePointers...); err != nil {
					log.Fatal(err)
				}

				rowData := make(map[string]string)
				for i, col := range columns {
					log.Printf("%v", values[i])
					rowData[col] = fmt.Sprintf("%v", values[i])
				}

				// Handle the row data as needed (e.g., print it)
				log.Println(rowData)
			}

			log.Println("finishedd")
		}
		return nil
	}
}
