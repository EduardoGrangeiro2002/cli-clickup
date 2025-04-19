package views

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListTaskTableModel struct {
	table table.Model
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m ListTaskTableModel) Init() tea.Cmd {
	return nil
}

func (m ListTaskTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ListTaskTableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func InitializeTaskTable() ListTaskTableModel {
	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Nome", Width: 15},
		{Title: "Status", Width: 15},
		{Title: "Data", Width: 15},
		{Title: "Prioridade", Width: 15},
		{Title: "Tempo estimado", Width: 15},
	}

	rows := []table.Row{
		{"1", "Tarefa 1", "Pendente", "2023-10-01", "Alta", "2h"},
		{"2", "Tarefa 2", "Em andamento", "2023-10-02", "Média", "1h"},
		{"3", "Tarefa 3", "Concluída", "2023-10-03", "Baixa", "30m"},
		{"4", "Tarefa 4", "Pendente", "2023-10-04", "Alta", "3h"},
		{"5", "Tarefa 5", "Em andamento", "2023-10-05", "Média", "1h"},
		{"6", "Tarefa 6", "Concluída", "2023-10-06", "Baixa", "30m"},
		{"7", "Tarefa 7", "Pendente", "2023-10-07", "Alta", "2h"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
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

	m := ListTaskTableModel{t}

	return m
}
