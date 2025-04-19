package components

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ComboBox struct {
	table table.Model
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m ComboBox) Init() tea.Cmd {
	return nil
}

func (m ComboBox) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(tea.Printf("Tarefa selecionada: %s", m.table.SelectedRow()[1]))
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m ComboBox) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func InitializeTaskTable(tableRows []table.Row) ComboBox {
	columns := []table.Column{
		{Title: "Options", Width: 5},
	}

	rows := tableRows

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(2),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := ComboBox{t}

	return m
}
