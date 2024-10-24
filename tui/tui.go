package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpxcz/sqlterm/tui/commands"
	databasesModel "github.com/jpxcz/sqlterm/tui/databases_panel"
	"github.com/jpxcz/sqlterm/tui/history_panel"
	queryModel "github.com/jpxcz/sqlterm/tui/query_panel"
	selectModel "github.com/jpxcz/sqlterm/tui/selection_panel"
)

type sessionState uint

type UiDimensions struct {
	height int
	width  int
}

type databaseConnected struct {
	key string
}

const (
	selectionView sessionState = iota
	queryView
	databasesView
	historyView
)

type mainModel struct {
	state                    sessionState
	selectDatabasePanelModel selectModel.SelectModel
	queryPanelModel          queryModel.QueryModel
	databasesPanelModel      databasesModel.DatabaseModel
	historyPanelModel        history_panel.HistoryModel
	uiDimensions             UiDimensions
}

func newModel() mainModel {
	m := mainModel{
		state: selectionView,
	}

	m.selectDatabasePanelModel = selectModel.NewSelectModel()
	m.queryPanelModel = queryModel.NewQueryModel()
	m.databasesPanelModel = databasesModel.NewDatabaseModel()
	m.historyPanelModel = history_panel.NewHistoryModel()
	m.uiDimensions = UiDimensions{height: 0, width: 0}

	return m
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == selectionView {
				m.state = queryView
			} else if m.state == queryView {
				m.state = databasesView
			} else if m.state == databasesView {
				m.state = historyView
			} else {
				m.state = selectionView
			}
		}
	case tea.WindowSizeMsg:
		m.uiDimensions.height = msg.Height
		m.uiDimensions.width = msg.Width
	case commands.MsgSyncConnectedDatabases:
		newDatabasePanelModel, newCmd := m.databasesPanelModel.Update(msg)
		databaseModel, ok := newDatabasePanelModel.(databasesModel.DatabaseModel)
		if !ok {
			panic("model is not a DatabaseModel")
		}
		m.databasesPanelModel = databaseModel
		cmd = newCmd
	case commands.MsgDatabaseQuery:
		newDatabasePanelModel, newCmd := m.databasesPanelModel.Update(msg)
		databaseModel, ok := newDatabasePanelModel.(databasesModel.DatabaseModel)
		if !ok {
			panic("model is not a DatabaseModel")
		}
		m.databasesPanelModel = databaseModel
		cmds = append(cmds, newCmd)

		newHistoryPanelMode, newCmd := m.historyPanelModel.Update(msg)
		historyModel, ok := newHistoryPanelMode.(history_panel.HistoryModel)
		if !ok {
			panic("model is not HistoryModel")
		}
		m.historyPanelModel = historyModel
		cmds = append(cmds, newCmd)
	case commands.MsgHistoryLookup:
		newDatabasePanelModel, newCmd := m.databasesPanelModel.Update(msg)
		databaseModel, ok := newDatabasePanelModel.(databasesModel.DatabaseModel)
		if !ok {
			panic("model is not a DatabaseModel")
		}
		m.databasesPanelModel = databaseModel
		cmds = append(cmds, newCmd)
	}

	switch m.state {
	case selectionView:
		newSelectModel, newCmd := m.selectDatabasePanelModel.Update(msg)
		selectionModel, ok := newSelectModel.(selectModel.SelectModel)
		if !ok {
			panic("model is not a SelectModel")
		}
		m.selectDatabasePanelModel = selectionModel
		cmd = newCmd
	case queryView:
		newQueryModel, newCmd := m.queryPanelModel.Update(msg)
		queryModel, ok := newQueryModel.(queryModel.QueryModel)
		if !ok {
			panic("model is not a QueryModel")
		}
		m.queryPanelModel = queryModel
		cmd = newCmd
	case databasesView:
		newDatabaseModel, newCmd := m.databasesPanelModel.Update(msg)
		databaseModel, ok := newDatabaseModel.(databasesModel.DatabaseModel)
		if !ok {
			panic("model is not a DatabaseModel")
		}
		m.databasesPanelModel = databaseModel
		cmd = newCmd
	case historyView:
		newHistoryPanelModel, newCmd := m.historyPanelModel.Update(msg)
		historyModel, ok := newHistoryPanelModel.(history_panel.HistoryModel)
		if !ok {
			panic("model is not a HistoryModel")
		}
		m.historyPanelModel = historyModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	if m.state == selectionView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			panelStyleFocused(15, m.uiDimensions.height-2).Render(
				m.selectDatabasePanelModel.View(),
			),
			lipgloss.JoinVertical(
				lipgloss.Left,
				panelStyleDefault(m.uiDimensions.width-25-6, 4-2).Render(m.queryPanelModel.View()),
				panelStyleDefault(
					m.uiDimensions.width-25-6,
					m.uiDimensions.height-8-2,
				).Render(m.databasesPanelModel.View()),
			),
			panelStyleDefault(10, m.uiDimensions.height-2).Render(
				m.historyPanelModel.View(),
			),
		)
	} else if m.state == queryView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			panelStyleDefault(15, m.uiDimensions.height-2).Render(
				m.selectDatabasePanelModel.View(),
			),
			lipgloss.JoinVertical(
				lipgloss.Left,
				panelStyleFocused(m.uiDimensions.width-25-6, 4-2).Render(m.queryPanelModel.View()),
				panelStyleDefault(m.uiDimensions.width-25-6, m.uiDimensions.height-8-2).Render(m.databasesPanelModel.View()),
			),
			panelStyleDefault(10, m.uiDimensions.height-2).Render(
				m.historyPanelModel.View(),
			),
		)
	} else if m.state == databasesView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			panelStyleDefault(15, m.uiDimensions.height-2).Render(
				m.selectDatabasePanelModel.View(),
			),
			lipgloss.JoinVertical(
				lipgloss.Left,
				panelStyleDefault(m.uiDimensions.width-25-6, 4-2).Render(m.queryPanelModel.View()),
				panelStyleFocused(m.uiDimensions.width-25-6, m.uiDimensions.height-8-2).Render(m.databasesPanelModel.View()),
			),
			panelStyleDefault(10, m.uiDimensions.height-2).Render(
				m.historyPanelModel.View(),
			),
		)
	} else if m.state == historyView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			panelStyleDefault(15, m.uiDimensions.height-2).Render(
				m.selectDatabasePanelModel.View(),
			),
			lipgloss.JoinVertical(
				lipgloss.Left,
				panelStyleDefault(m.uiDimensions.width-25-6, 4-2).Render(m.queryPanelModel.View()),
				panelStyleDefault(m.uiDimensions.width-25-6, m.uiDimensions.height-8-2).Render(m.databasesPanelModel.View()),
			),
			panelStyleFocused(10, m.uiDimensions.height-2).Render(
				m.historyPanelModel.View(),
			),
		)
	}
	return s
}

func NewTeaProgram() *tea.Program {
	p := tea.NewProgram(newModel())
	return p
}
