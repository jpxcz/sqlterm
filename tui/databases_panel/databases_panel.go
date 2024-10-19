package databases_panel

import (
	"database/sql"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/tui/commands"
)

type ConnectedDatabase struct {
	db    *sql.DB
	value string
}

type DatabaseModel struct {
	databases map[string]*ConnectedDatabase
	value     string
}

func (m *DatabaseModel) updateCurrentDatabases() {
	dbs := databases.GetDatabases()

	for key, db := range dbs {
		if m.databases[key] != nil && db.ConnectionStatus != databases.DbConnected {
            delete(m.databases, key)
		} else if m.databases[key] == nil && db.ConnectionStatus == databases.DbConnected {
			m.databases[key] = &ConnectedDatabase{db: db.Db, value: key}
		}
	}
}

func NewDatabaseModel(value string) DatabaseModel {
	return DatabaseModel{
		databases: make(map[string]*ConnectedDatabase),
		value:     value,
	}
}

func (m DatabaseModel) Init() tea.Cmd {
	return nil
}

func (m DatabaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			log.Println("tab pressed")
		}
	case commands.MsgDatabaseSelectionUpdate:
		m.updateCurrentDatabases()
	}

	return m, nil
}

func (m DatabaseModel) View() string {
	s := "Databases\n"
	for key := range m.databases {
		s += key + "\n"
	}
	return s
}
