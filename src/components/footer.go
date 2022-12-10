package footer

import (
	"strings"

	"github.com/2mer/goplaylist/sounds"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	PlaySound key.Binding
	Quit      key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.PlaySound, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PlaySound, k.Quit},
	}
}

var keys = keyMap{
	PlaySound: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "play sound"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type Model struct {
	keys       keyMap
	help       help.Model
	inputStyle lipgloss.Style
	lastKey    string
	Quitting   bool
}

func NewModel() Model {
	return Model{
		keys:       keys,
		help:       help.New(),
		inputStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PlaySound):
			sounds.PlaySound("gunag")
		case key.Matches(msg, m.keys.Quit):
			m.Quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {

	helpView := m.help.View(m.keys)
	height := 8 - strings.Count(helpView, "\n")

	return "\n" + strings.Repeat("\n", height) + helpView
}
