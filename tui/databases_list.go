package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle()

type item struct {
	title string
	value string
}

type databaseModelList struct {
	list list.Model
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.value
}

func (i item) FilterValue() string {
	return i.title
}

func (m databaseModelList) Init() tea.Cmd {
    return nil;
}

func (m databaseModelList) View() string {
	return docStyle.Render(m.list.View())
}

func (m databaseModelList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func NewDatabasesModelList() tea.Model {
	items := []list.Item{
		item{"DB1", "mysql -u root ,,,,"},
		item{"DB2", "mysql -u root ,,,,"},
		item{"DB3", "mysql -u root ,,,,"},
		item{"DB4", "mysql -u root ,,,,"},
	}

	m := databaseModelList{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}

    m.list.Title = "Databases"
    return m
}
