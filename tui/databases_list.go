package tui

import (
	// "fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Choice struct {
	title    string
	value    string
	selected bool
}

func (c Choice) View(isCursor bool) string {
	var text string

	if c.selected {
		text = "[x] " + c.title
	} else {
		text = "[ ] " + c.title
	}

	if isCursor {
		return lipgloss.NewStyle().Bold(true).Render(text)
	}

	return text
}

type SelectModel struct {
	cursor  int
	choices []Choice
}

func NewSelectModel() SelectModel {
	return SelectModel{
		cursor: 0,
		choices: []Choice{
			{"DB1", "mysql -u root ,,,,", false},
			{"DB2", "mysql -u root ,,,,", false},
			{"DB3", "mysql -u root ,,,,", false},
		},
	}
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.toogleOption()
		case "j", "down":
			m.cursor = nextElement(m)
		case "k", "up":
			m.cursor = previousElement(m)
		}
	}

	return m, nil
}

func (m SelectModel) View() string {
	s := strings.Builder{}
	for i, c := range m.choices {
		s.WriteString(c.View(i == m.cursor))
		s.WriteString("\n")
	}

	return s.String()
}

func (m SelectModel) toogleOption() {
	if m.cursor >= len(m.choices) || m.cursor < 0 {
		return
	}

	m.choices[m.cursor].selected = !m.choices[m.cursor].selected
}

func nextElement(m SelectModel) int {
	m.cursor++
	if m.cursor >= len(m.choices) {
		m.cursor = 0
	}

    return m.cursor 
}

func  previousElement(m SelectModel) int {
	m.cursor--
	if m.cursor < 0 {
		m.cursor = len(m.choices) - 1
	}

    return m.cursor
}

// var docStyle = lipgloss.NewStyle()
//
// type item struct {
// 	title string
// 	value string
// }
//
// type databaseOption struct {
// 	title        string
// 	connectivity string
// }
//
// type databasesModel struct {
// 	cursor          uint
// 	choices         []databaseOption
// 	selectedChoices []uint
// }
//
// type databaseModelList struct {
// 	list list.Model
// }
//
// func (i item) Title() string {
// 	return i.title
// }
//
// func (i item) Description() string {
// 	return i.value
// }
//
// func (i item) FilterValue() string {
// 	return i.title
// }
//
// func (m databaseModelList) Init() tea.Cmd {
// 	return nil
// }
//
// func (m databaseModelList) View() string {
// 	return docStyle.Render(m.list.View())
// }
//
// func (m databaseModelList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }
//
// func NewDatabasesModelList() tea.Model {
// 	items := []list.Item{
// 		item{"DB1", "mysql -u root ,,,,"},
// 		item{"DB2", "mysql -u root ,,,,"},
// 		item{"DB3", "mysql -u root ,,,,"},
// 		item{"DB4", "mysql -u root ,,,,"},
// 	}
//
// 	m := databaseModelList{
// 		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
// 	}
//
// 	m.list.Title = "Databases"
// 	return m
// }
