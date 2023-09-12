package main

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type item struct {
	name string
	done bool
}

type model struct {
	table table.Model
	input textinput.Model
	items []*item
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) rows() (rows []table.Row) {
	for _, item := range m.items {
		var done string

		if item.done {
			done = "✅"
		}

		rows = append(rows, table.Row{done, item.name})
	}

	return
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.input.Focused() {
			switch msg.String() {
			case "enter":
				m.items = append(m.items, &item{m.input.Value(), false})
				m.table.SetRows(m.rows())
				m.input.SetValue("")
				m.input.Blur()
				return m, nil
			case "esc":
				m.input.SetValue("")
				m.input.Blur()
				return m, nil
			default:
				m.input, cmd = m.input.Update(msg)
				return m, cmd
			}
		}

		switch msg.String() {
		case " ":
			item := m.items[m.table.Cursor()]
			item.done = !item.done
			m.table.SetRows(m.rows())
			return m, nil
		case "backspace":
			if len(m.items) > 0 {
				idx := m.table.Cursor()
				m.items = append(m.items[:idx], m.items[idx+1:]...)
				m.table.SetRows(m.rows())

				if idx > len(m.items)-1 {
					m.table.SetCursor(idx - 1)
				}
			}

			return m, nil
		case "enter":
			m.input.Focus()
			return m, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.input.View() + "\n" + baseStyle.Render(m.table.View())
}

func main() {
	columns := []table.Column{
		{Width: 2},
		{Width: 50},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	i := textinput.New()
	i.Placeholder = "New todo..."

	m := model{t, i, []*item{}}

	m.table.SetRows(m.rows())

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
