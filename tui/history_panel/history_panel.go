package history_panel

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpxcz/sqlterm/nodequery"
	"github.com/jpxcz/sqlterm/tui/commands"
)

func view(n int, isCursor bool) string {
	text := strconv.Itoa(n)

	if isCursor {
		return lipgloss.NewStyle().
			Bold(true).
			Render(text)
	}

	return lipgloss.NewStyle().
		Render(text)
}

type HistoryModel struct {
	cursor     int
	maxChoices int
	query      string
	dbKey      string
}

func NewHistoryModel() HistoryModel {
	return HistoryModel{
		cursor:     0,
		maxChoices: 0,
	}
}

func (m HistoryModel) Init() tea.Cmd {
	return nil
}

func (m HistoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			m.nextElement()
			cmd = commands.CmdHistoryLookup(m.cursor)
		case "k", "up":
			m.previousElement()
			cmd = commands.CmdHistoryLookup(m.cursor)
		}

	case commands.MsgDatabaseQuery:
		newQuery := string(msg)
		if m.query != newQuery {
			m.cursor = 0
		}
		m.query = string(msg)
		m.getMaxAmountOfChoices()
        cmd = commands.CmdHistoryLookup(m.cursor)
	}

	return m, cmd
}

func (m *HistoryModel) getMaxAmountOfChoices() {
	n := nodequery.GetMaxLenghtOfQueryNodes(m.query)
	m.maxChoices = n
}

func (m HistoryModel) View() string {
	s := strings.Builder{}
	s.WriteString("History\n")
	for i := 0; i <= m.maxChoices; i++ {
		s.WriteString(view(i, i == m.cursor))
		s.WriteString("\n")
	}
	return s.String()
}

func (m *HistoryModel) nextElement() {
	m.cursor++
	if m.cursor > m.maxChoices {
		m.cursor = 0
	}
}

func (m *HistoryModel) previousElement() {
	m.cursor--
	if m.cursor < 0 {
		m.cursor = m.maxChoices
	}
}
