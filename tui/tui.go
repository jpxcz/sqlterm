package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState uint

const (
	dbSelectorView sessionState = iota
	dbView
)

var (
	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(40).
			Align(lipgloss.Center)
	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(40).
				Align(lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#FF00FF"))
)

type mainModel struct {
	state     sessionState
	index     int
	databases tea.Model
	database  tea.Model // todo: change to the right tabs
}

func newModel() mainModel {
	m := mainModel{
		state: dbSelectorView,
	}
    m.databases = NewDatabasesModelList()
    m.database = NewDatabasesModelList()

	return m
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == dbSelectorView {
				m.state = dbView
			} else {
				m.state = dbSelectorView
			}
		}
	}
	return m, cmd
}

func (m mainModel) View() string {
	var s string
	if m.state == dbSelectorView {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			focusedModelStyle.Render(
				modelStyle.Render(m.databases.View()),
			),
			modelStyle.Render(m.database.View()),
		)
	} else {
		s += lipgloss.JoinHorizontal(
			lipgloss.Top,
			modelStyle.Render(
				modelStyle.Render(m.databases.View()),
			),
			focusedModelStyle.Render(m.database.View()),
		)
	}
	return s
}

func NewTeaProgram() *tea.Program {
    p := tea.NewProgram(newModel())
    return p
}
