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
	value            string
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

func NewSelectModel() SelectModel {
	return SelectModel{
		cursor: 0,
		choices: []choice{
			{"DB1", "DB1", false, 0},
			{"DB2", "DB2", false, 0},
			{"DB3", "DB3", false, 0},
		},
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

	if !m.choices[m.cursor].selected {
		db, err := databases.ConnectToDatabase(
			"DB1",
			"user1",
			"0.0.0.0",
			"3306",
			"password1",
			"mysql",
		)
        if err != nil {
            log.Println("error connecting to database")
        }

        log.Println("connected to database")

		m.choices[m.cursor].connectionStatus = uint(db.ConnectionStatus)
		m.choices[m.cursor].selected = true
		return commands.CmdDatabaseSelectionUpdate
	}

	databases.DisconnectFromDatabase("DB1")
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
