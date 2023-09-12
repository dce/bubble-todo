package main

import (
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	name string
	done bool
}

type model struct {
	table table.Model
	items []item
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
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.table.View()
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

	m := model{t, []item{
		{"One", false},
		{"Two", true},
		{"Three", false},
	}}

	m.table.SetRows(m.rows())

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error running program:", err)
	}

	log.Println("Hello!")
}
