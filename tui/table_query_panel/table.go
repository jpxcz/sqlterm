package table_query_panel

import (
	"log"

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
	query            string
	lookupHistory    int
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case commands.MsgDatabasePanelSelectionUpdate:
		m.selectedDatabase = string(msg)
		node := nodequery.GetQueryNode(m.query, m.selectedDatabase, m.lookupHistory)
		m.node = node
	case commands.MsgDatabaseQuery:
		if m.selectedDatabase == "" {
			break
		}

		m.query = string(msg)
		m.lookupHistory = 0
		node := nodequery.GetQueryNode(m.query, m.selectedDatabase, 0)
		m.node = node
	case commands.MsgHistoryLookup:
		lookupHistory := int(msg)
		log.Println("lookupHistory", lookupHistory)
		m.lookupHistory = lookupHistory
		node := nodequery.GetQueryNode(m.query, m.selectedDatabase, m.lookupHistory)

		log.Println("node", node)
		m.node = node
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	if m.node == nil {
		return ""
	}

	columns := m.node.GetColumnsViewData()
	rows := m.node.GetRowsViewData()

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
