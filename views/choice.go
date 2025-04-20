package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const dotChar = " • "

var (
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
)

type ChoiceModel struct {
	Choice int
	Chosen bool
}

func (m ChoiceModel) Init() tea.Cmd {
	m.Choice = -1
	return nil
}

func (m ChoiceModel) View() string {
	var s string

	if !m.Chosen {
		s = choicesView(m)
	}

	return mainStyle.Render("\n" + s + "\n\n")
}

func (m ChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.Chosen {
		return updateChoices(msg, m)
	}
	return m, nil
}

func updateChoices(msg tea.Msg, m ChoiceModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < -1 {
				m.Choice = -1
			}
		case "enter":
			m.Chosen = true
			return m, nil
		}
	}

	return m, nil
}

func choicesView(m ChoiceModel) string {
	c := m.Choice

	tpl := "Bem vindo a CLI de gerenciamento de projetos\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Criar uma nova tarefa", c == 0),
		checkbox("Listar todas as tarefas", c == 1),
		checkbox("Criar uma nova reunião", c == 2),
		checkbox("Listas todas as reuniões", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
