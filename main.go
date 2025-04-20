package main

import (
	"fmt"

	"github.com/EduardoGrangeiro2002/cli-clikup/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

// General stuff for styling the view
var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

type AppModel struct {
	currentView   int
	choice        views.ChoiceModel
	taskForm      *views.CreateTaskFormModel
	listTaskTable views.ListTaskTableModel
	// createMeetingForm views.CeateMeetingFormModel
	// listMeetingTable  views.ListMeetingTableModel
	Quitting bool
}

func initialModel() AppModel {
	return AppModel{
		currentView:   -1,
		choice:        views.ChoiceModel{},
		taskForm:      views.NewModel(),
		listTaskTable: views.InitializeTaskTable(),
		// createMeetingForm: CeateMeetingFormModel{},
		// listMeetingTable:  ListMeetingTableModel{},
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()

		if k == "ctrl+r" {
			m.currentView = -1
			m.choice.Choice = -1
			m.choice.Chosen = false
			m.taskForm = views.NewModel()
		}

		if k == "esc" || k == "ctrl+c" {
			if m.currentView != -1 {
				switch m.currentView {
				case 0:
					taskFormModel, cmd := m.taskForm.Update(msg)
					m.taskForm = taskFormModel.(*views.CreateTaskFormModel)
					return m, cmd
				}
			}
			m.Quitting = true
			return m, tea.Quit
		}
	}

	switch m.currentView {
	case -1:
		choiceModel, choiceCmd := m.choice.Update(msg)
		m.choice = choiceModel.(views.ChoiceModel)
		cmds = append(cmds, choiceCmd)
	case 0:
		taskFormModel, taskFormCmd := m.taskForm.Update(msg)
		m.taskForm = taskFormModel.(*views.CreateTaskFormModel)
		cmds = append(cmds, taskFormCmd)
	case 1:
		listTaskTableModel, listTaskTableCmd := m.listTaskTable.Update(msg)
		m.listTaskTable = listTaskTableModel.(views.ListTaskTableModel)
		cmds = append(cmds, listTaskTableCmd)
		//
		// case "createMeetingForm":
		// 	createMeetingFormModel, createMeetingFormCmd := m.createMeetingForm.Update(msg)
		// 	m.createMeetingForm = createMeetingFormModel.(CeateMeetingFormModel)
		// 	cmds = append(cmds, createMeetingFormCmd)
		//
		// case "listMeetingTable":
		// 	listMeetingTableModel, listMeetingTableCmd := m.listMeetingTable.Update(msg)
		// 	m.listMeetingTable = listMeetingTableModel.(ListMeetingTableModel)
		// 	cmds = append(cmds, listMeetingTableCmd)
	}
	if m.choice.Chosen {
		m.currentView = m.choice.Choice
	}

	if m.currentView == 0 && m.taskForm.SendData {
		m.currentView = -1
		m.choice.Choice = -1
		m.choice.Chosen = false
		m.taskForm = views.NewModel()
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	var content string
	switch m.currentView {
	case -1:
		content = m.choice.View()
	case 0:
		content = m.taskForm.View()
	case 1:
		content = m.listTaskTable.View()
		// case "createMeetingForm":
		// 	content = m.createMeetingForm.View()
		// case "listMeetingTable":
		// 	content = m.listMeetingTable.View()
	}

	return content + "\n\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
