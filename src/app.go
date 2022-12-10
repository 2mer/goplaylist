package main

import (
	"fmt"

	footer "github.com/2mer/goplaylist/components"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type model struct {
	spinner spinner.Model
	footer  footer.Model
	err     error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	f := footer.NewModel()

	return model{spinner: s, footer: f}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		if cmd == nil {
			m.footer, cmd = m.footer.Update(msg)
		}

		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf(`
	
%s Loading forever...

%s
`, m.spinner.View(), m.footer.View())
	if m.footer.Quitting {
		return str + "Quitting... \n"
	}
	return str
}
