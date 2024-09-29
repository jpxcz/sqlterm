package selection_panel

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type choice struct {
	title    string
	value    string
	selected bool
}

func (c choice) view(isCursor bool) string {
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
	choices []choice
}

func NewSelectModel() SelectModel {
	return SelectModel{
		cursor: 0,
		choices: []choice{
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
			m.nextElement()
		case "k", "up":
			m.previousElement()
		}
	}

	return m, nil
}

func (m SelectModel) View() string {
	s := strings.Builder{}
    s.WriteString("Databases\n")
	for i, c := range m.choices {
		s.WriteString(c.view(i == m.cursor))
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

func (m *SelectModel) nextElement() {
	m.cursor++
	if m.cursor >= len(m.choices) {
		m.cursor = 0
	}
}

func (m *SelectModel) previousElement() {
	m.cursor--
	if m.cursor < 0 {
		m.cursor = len(m.choices) - 1
	}

}

