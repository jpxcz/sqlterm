package databases_panel

import (
	tea "github.com/charmbracelet/bubbletea"
)

type DatabaseModel struct {
    value string
}

func NewDatabaseModel(value string) DatabaseModel {
    return DatabaseModel{value: value}
}

func (m DatabaseModel) Init() tea.Cmd {
    return nil
}

func (m DatabaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil
}

func (m DatabaseModel) View() string {
    return m.value
}
