package commands

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/nodequery"
)

type MsgSyncConnectedDatabases bool

type MsgDatabaseQuery string

type MsgDatabasePanelSelectionUpdate string

type MsgHistoryLookup int

func CmdDatabaseSelectionUpdate() tea.Msg {
	return MsgSyncConnectedDatabases(true)
}

func CmdDatabaseQuery(query string) tea.Cmd {
	return func() tea.Msg {
		dbs := databases.GetDatabases()
		for key, db := range dbs {
			if db.ConnectionStatus != databases.DbConnected {
				continue
			}

			rows, err := db.Query(query)
			if err != nil {
				log.Println("could not run query", query, err)
				continue
			}

			result, err := nodequery.NewQueryResult(query, rows)
			if err != nil {
				log.Println("problem parsing query result rows", query, err)
				continue
			}

			nodequery.AttachQueryResult(query, key, result)

		}

		return MsgDatabaseQuery(query)
	}
}

func CmdHistoryLookup(n int) tea.Cmd {
    return func() tea.Msg {
        return MsgHistoryLookup(n)
    }
}
