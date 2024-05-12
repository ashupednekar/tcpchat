package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	quitting    bool
	title       string
	description string
}

func (m Model) Title() string       { return m.title }
func (m Model) Description() string { return m.description }
func (m Model) FilterValue() string { return m.title }

func NewModel() Model {
	return Model{
		title:       "chat",
		description: "chat",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	return "coming soon"
}
