package table_query_panel

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpxcz/sqlterm/nodequery"
	"github.com/jpxcz/sqlterm/tui/commands"
)

func NewTableModel(selectedDatabase string) TableModel {
	return TableModel{
		selectedDatabase: selectedDatabase,
	}
}

type TableModel struct {
	table            table.Model
	selectedDatabase string
	node             *nodequery.QueryResult
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// var cmds []tea.Cmd

	switch msg := msg.(type) {
	case commands.MsgDatabasePanelSelectionUpdate:
		m.selectedDatabase = string(msg)
	case commands.MsgDatabaseQuery:
		if m.selectedDatabase == "" {
			break
		}

		node := nodequery.GetLastQueryForDatabase(m.selectedDatabase)
		if node == nil {
			break
		}
		m.node = node
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	node := nodequery.GetLastQueryForDatabase(m.selectedDatabase)
	if node == nil {
		return ""
	}

	columns := node.GetColumnsViewData()
	rows := node.GetRowsViewData()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t.View()
}
