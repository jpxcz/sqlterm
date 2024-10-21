package databases_panel

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/tui/commands"
)

type ConnectedDatabase struct {
	database *databases.Database
	value    string // will be changed for the query results 
}

type DatabaseModel struct {
	databases map[string]*ConnectedDatabase
	activeTab int
}

func (m *DatabaseModel) updateCurrentDatabases() {
	dbs := databases.GetDatabases()

	for key, db := range dbs {
		if m.databases[key] != nil && db.ConnectionStatus != databases.DbConnected {
			delete(m.databases, key)
		} else if m.databases[key] == nil && db.ConnectionStatus == databases.DbConnected {
            m.databases[key] = &ConnectedDatabase{database: db, value: key}
		}
	}
}

func NewDatabaseModel(value string) DatabaseModel {
	return DatabaseModel{
		databases: make(map[string]*ConnectedDatabase),
		activeTab: 1,
	}
}

func (m DatabaseModel) Init() tea.Cmd {
	return nil
}

func (m DatabaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		i, err := strconv.Atoi(msg.String())
		if err != nil {
			break
		}

		if i <= len(m.databases) {
			m.activeTab = i
		} else {
			m.activeTab = 1
		}

	case commands.MsgDatabaseSelectionUpdate:
		m.updateCurrentDatabases()
	}

	return m, nil
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				BorderForeground(highlightColor).
				Padding(0, 1)
	activeTabStyle = inactiveTabStyle.Border(activeTabBorder, true)
	windowStyle    = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(2, 0).
			Align(lipgloss.Center).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop()
)

func (m DatabaseModel) View() string {
	doc := strings.Builder{}
	var renderedTabs []string

	tabs := make([]string, 0, len(m.databases))
	for key := range m.databases {
		tabs = append(tabs, key)
	}

	if len(tabs) == 0 {
		return "not connected yet"
	}

	activeTab := m.activeTab - 1
	activeKeyTab := tabs[activeTab]

	for i, t := range tabs {
		var style lipgloss.Style
		isFirst := i == 0
		isLast := i == len(tabs)-1
		isActive := i == activeTab

		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}

		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}

		style = style.Border(border)
		text := "(" + strconv.Itoa(i+1) + ") " + t
		renderedTabs = append(renderedTabs, style.Render(text))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(
		windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).
			Render(m.databases[activeKeyTab].value),
	)
	return docStyle.Render(doc.String())

}
