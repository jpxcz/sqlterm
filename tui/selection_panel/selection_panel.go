package selection_panel

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/tui/commands"
)

type choice struct {
	title            string
	selected         bool
	connectionStatus uint
}

func (c choice) getConnectionColor() string {
	switch c.connectionStatus {
	case 1:
		return "#00ff00"
	case 2:
		return "#ff0000"
	default:
		return "#ffffff"
	}
}

func (c choice) view(isCursor bool) string {
	var text string

	if c.selected {
		text = "[x] " + c.title
	} else {
		text = "[ ] " + c.title
	}

	if isCursor {
		return lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(c.getConnectionColor())).
			Render(text)
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(c.getConnectionColor())).
		Render(text)
}

type SelectModel struct {
	cursor  int
	choices []choice
}

func newChoices(dbs map[string]*databases.Database) []choice {
	choices := make([]choice, 0)

	for key, db := range dbs {
		choices = append(choices, choice{
			title:            key,
			selected:         false,
			connectionStatus: uint(db.ConnectionStatus),
		})
	}

	return choices
}

func NewSelectModel() SelectModel {
	dbs := databases.GetDatabases()
	choices := newChoices(dbs)

	return SelectModel{
		cursor:  0,
		choices: choices,
	}
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			cmd = m.toogleOption()
		case "j", "down":
			m.nextElement()
		case "k", "up":
			m.previousElement()
		}
	}

	return m, cmd
}

func (m SelectModel) View() string {
	s := strings.Builder{}
	s.WriteString("Databases\n")
	for i, c := range m.choices {
		s.WriteString(c.view(i == m.cursor))
		s.WriteString("\n")
	}

	return s.String()
}

func (m *SelectModel) toogleOption() tea.Cmd {
	if m.cursor >= len(m.choices) || m.cursor < 0 {
		return nil
	}

	keyOption := m.choices[m.cursor].title

	// To connect
	if !m.choices[m.cursor].selected {
        log.Println("connecting to choice", keyOption)
		db := databases.GetConnection(keyOption)

		dbConnectionStatus, err := databases.ConnectToDatabase(
			db.DatabaseCredentials.ShortName,
			db.DatabaseCredentials.Username,
            db.DatabaseCredentials.DatabaseName,
			db.DatabaseCredentials.Hostname,
			db.DatabaseCredentials.Port,
			db.DatabaseCredentials.Password,
			db.DatabaseCredentials.Type,
		)

		if err != nil {
			log.Println("cannot connect to choice", keyOption, err)
		}

		m.choices[m.cursor].connectionStatus = uint(dbConnectionStatus)
		m.choices[m.cursor].selected = true
		return commands.CmdDatabaseSelectionUpdate
	}

	// To disconnect
    log.Println("disconnecting from choice", keyOption)
	databases.DisconnectFromDatabase(keyOption)
	m.choices[m.cursor].connectionStatus = uint(databases.DbDisconnected)
	m.choices[m.cursor].selected = false

	return commands.CmdDatabaseSelectionUpdate
}

func (m *SelectModel) nextElement() {
	m.cursor++
	if m.cursor >= len(m.choices) {
		m.cursor = 0
	}
}

func (m *SelectModel) previousElement() {
	m.cursor--
	if m.cursor < 0 {
		m.cursor = len(m.choices) - 1
	}
}
