package views

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type CreateTaskFormModel struct {
	form          *huh.Form
	name          string
	description   string
	status        string
	owners        []string
	priority      string
	date          string
	stimatedTime  string
	confirmedForm bool
	quitting      bool
	SendData      bool
}

func NewModel() *CreateTaskFormModel {
	m := &CreateTaskFormModel{}
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nome").
				Value(&m.name).
				Validate(nameValidator),
			huh.NewText().
				Title("Descricão").
				CharLimit(900).
				Value(&m.description).
				Validate(descriptionValidator),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Pendente", "Concluido", "Cancelado")...).
				Title("Status").
				Value(&m.status),
			huh.NewMultiSelect[string]().
				Options(huh.NewOptions("Jefferson", "Robson", "Zolk", "Eduardo", "Luciano")...).
				Title("Responsáveis").
				Value(&m.owners),
			huh.NewInput().
				Title("Data").
				Value(&m.date).
				Validate(dateValidator).
				CharLimit(36),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Urgente", "Alta", "Normal", "Baixa")...).
				Title("Prioridade").
				Value(&m.priority),

			huh.NewSelect[string]().
				Options(huh.NewOptions("1h", "2h", "4h", "6h", "8h", "12", "24")...).
				Title("Tempo estimado").
				Value(&m.stimatedTime),
			huh.NewConfirm().
				Title("Criar tarefa?").
				Value(&m.confirmedForm),
		),
	)

	m.form.NextField()
	m.form.PrevField()

	return m
}

func (m *CreateTaskFormModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m *CreateTaskFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}
	form, cmd := m.form.Update(msg)

	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m *CreateTaskFormModel) View() string {
	if m.quitting {
		return "Formulário cancelado"
	}

	if m.form.State == huh.StateCompleted && m.confirmedForm {
		_ = spinner.New().Title("Enviando a tarefa para o ClickUp").Action(sendData).Run()
		m.SendData = true
		return "Tarefa enviada com sucesso!\nPressione qualquer tecla para continuar..."
	}
	return m.form.View()
}

func sendData() {
	time.Sleep(2 * time.Second)
}

func dateValidator(input string) error {
	if (len(input) == 0 || len(input) > 36) && (len(input) < 10 || len(input) > 16) {
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
