package tui

import (
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModelTui(t *testing.T) {
	m := NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		t.Fatal(err)
		os.Exit(1)
	}
}
