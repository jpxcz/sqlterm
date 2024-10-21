package commands

import (
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

			rows, err := db.Db.Query(query)
			if err != nil {
				log.Println(
					"Error executing query for database",
					db.DatabaseCredentials.ShortName,
					err,
				)
                continue
			}

            log.Println(rows)
		}

        return nil
	}
}
