package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	name = iota
	description
	status
	date
	owners
	priority
	stimatedTime
)

type item string

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type CreateTaskFormModel struct {
	inputs          []textinput.Model
	focused         int
	statusOptions   []string
	showDropdown    bool
	dropdownFocused bool
	err             error
}

func dateValidator(input string) error {
	if (len(input) == 0 || len(input) > 35) && (len(input) < 10 || len(input) > 16) {
		return fmt.Errorf("a data da tarefa não pode ser vazia e deve ter o formato DD-MM-YYYY:hh:mm")
	}
	return nil
}

func nameValidator(input string) error {
	if len(input) == 0 || len(input) > 45 {
		return fmt.Errorf("o nome da tarefa não pode ser vazio")
	}
	return nil
}

func descriptionValidator(input string) error {
	if len(input) == 0 || len(input) > 900 {
		return fmt.Errorf("a descrição da tarefa não pode ser vazia")
	}
	return nil
}

func statusValidator(input string) error {
	if len(input) == 0 || len(input) > 35 {
		return fmt.Errorf("o status da tarefa não pode ser vazio")
	}
	return nil
}

func ownersValidator(input string) error {
	if len(input) == 0 || len(input) > 200 {
		return fmt.Errorf("os responsáveis da tarefa não podem ser vazios")
	}
	return nil
}

func priorityValidator(input string) error {
	if len(input) == 0 || len(input) > 35 {
		return fmt.Errorf("a prioridade da tarefa não pode ser vazia")
	}
	return nil
}

func estimatedTimeValidator(input string) error {
	if len(input) == 0 || len(input) > 20 {
		return fmt.Errorf("o tempo estimado da tarefa não pode ser vazio")
	}
	return nil
}

func InitialCreateTaskFormModel() CreateTaskFormModel {
	var inputs []textinput.Model = make([]textinput.Model, 7)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Digite o nome da tarefa"
	inputs[name].Focus()
	inputs[name].CharLimit = 45
	inputs[name].Width = 55
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "Digite a descrição da tarefa"
	inputs[description].CharLimit = 900
	inputs[description].Width = 1000
	inputs[description].Prompt = ""
	inputs[description].Validate = descriptionValidator

	inputs[status] = textinput.New()
	inputs[status].Placeholder = "Selecione o status da tarefa"
	inputs[status].CharLimit = 35
	inputs[status].Width = 45
	inputs[status].ShowSuggestions = true
	inputs[status].SetSuggestions([]string{"Pendente", "Concluída", "Cancelada"})
	inputs[status].Validate = statusValidator

	inputs[date] = textinput.New()
	inputs[date].Placeholder = "Digite a data da tarefa no formato DD-MM-YYYY:hh:mm"
	inputs[date].CharLimit = 35
	inputs[date].Width = 45
	inputs[date].Prompt = ""
	inputs[date].Validate = dateValidator

	inputs[owners] = textinput.New()
	inputs[owners].Placeholder = "Digite os responsáveis da tarefa no formato Name1,Name2"
	inputs[owners].CharLimit = 200
	inputs[owners].Width = 220
	inputs[owners].Prompt = ""
	inputs[owners].Validate = ownersValidator

	inputs[priority] = textinput.New()
	inputs[priority].Placeholder = "Digite a prioridade da tarefa"
	inputs[priority].CharLimit = 35
	inputs[priority].Width = 45
	inputs[priority].Prompt = ""
	inputs[priority].Validate = priorityValidator

	inputs[stimatedTime] = textinput.New()
	inputs[stimatedTime].Placeholder = "Digite o tempo estimado da tarefa no formato 3h30m"
	inputs[stimatedTime].CharLimit = 50
	inputs[stimatedTime].Width = 60
	inputs[stimatedTime].Prompt = ""
	inputs[stimatedTime].Validate = estimatedTimeValidator

	return CreateTaskFormModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
		statusOptions: []string{
			"Pendente",
			"Concluída",
			"Cancelada",
		},
		showDropdown: false,
	}
}

func (m CreateTaskFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateTaskFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m CreateTaskFormModel) View() string {
	return fmt.Sprintf(
		`Formulário para criar uma tarefa:

 %s
 %s

 %s
 %s
 
 %s  %s

 %s  %s

 %s  %s

 %s  %s
 %s  %s
 %s
`,
		inputStyle.Width(20).Render("Nome"),
		m.inputs[name].View(),
		inputStyle.Width(15).Render("Descrição"),
		m.inputs[description].View(),
		inputStyle.Width(10).Render("Status"),
		m.inputs[status].View(),
		inputStyle.Width(10).Render("Data"),
		m.inputs[date].View(),
		inputStyle.Width(15).Render("Responsáveis"),
		m.inputs[owners].View(),
		inputStyle.Width(15).Render("Prioridade"),
		m.inputs[priority].View(),
		inputStyle.Width(15).Render("Tempo estimado"),
		m.inputs[stimatedTime].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func (m *CreateTaskFormModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *CreateTaskFormModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
