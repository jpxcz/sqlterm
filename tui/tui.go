package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	styles "github.com/jpxcz/sqlterm/tui/styles"
)

type sessionState uint

const (
	selectionView sessionState = iota
	visualizerView
)

type mainModel struct {
	state     sessionState
	databases tea.Model
	database  DatabaseModel // todo: change to the right tabs
	height    int
}

func newModel() mainModel {
	m := mainModel{
		state: selectionView,
	}
	m.databases = NewSelectModel()
	m.database = NewDatabaseModel("fake database selected")

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
				m.state = visualizerView
			} else {
				m.state = selectionView
			}

		}
	case tea.WindowSizeMsg:
		m.height = msg.Height - 2
	}

	switch m.state {
	case selectionView:
		newSelectModel, newCmd := m.databases.Update(msg)
		selectionModel, ok := newSelectModel.(SelectModel)
		if !ok {
			panic("model is not a SelectModel")
		}
		m.databases = selectionModel
		cmd = newCmd
	case visualizerView:

	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	if m.state == selectionView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.DatabasesConnectivityViewStyleFocused(15, m.height).Render(
				m.databases.View(),
			),
			styles.DatabasesConnectivityViewStyleDefault(15, m.height).Render(
				m.database.View(),
			),
		)
	} else {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.DatabasesConnectivityViewStyleDefault(15, m.height).Render(
				m.databases.View(),
			),
			styles.DatabasesConnectivityViewStyleFocused(15, m.height).Render(
				m.database.View(),
			),
		)
	}
	return s
}

func NewTeaProgram() *tea.Program {
	p := tea.NewProgram(newModel())
	return p
}
