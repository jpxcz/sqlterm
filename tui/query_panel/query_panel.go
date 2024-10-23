package query_panel

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpxcz/sqlterm/tui/commands"
)

type QueryModel struct {
	textarea textarea.Model
}

func NewQueryModel() QueryModel {
	ta := textarea.New()
	ta.Placeholder = "Enter your query here"

	return QueryModel{
		textarea: ta,
	}
}

func (m QueryModel) Init() tea.Cmd {
	return nil
}

func (m QueryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+g":
			cmd = commands.CmdDatabaseQuery(m.textarea.Value())
			cmds = append(cmds, cmd)
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m QueryModel) View() string {
	return m.textarea.View()
}
