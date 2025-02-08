package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Board struct {
	help        help.Model
	loaded      bool
	focused     status // tracks selected column
	cols        []column
	quitting    bool
	statusMsg   string
	clearMsgCmd tea.Cmd // command to clear message after some time
}

type clearMessage struct{}
type taskMovedMessage struct{}

func NewBoard() *Board {
	help := help.New()
	help.ShowAll = true
	return &Board{help: help, focused: done}
}

func (m *Board) Init() tea.Cmd {
	return nil
}

func (m *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		m.help.Width = msg.Width - margin
		for i := 0; i < len(m.cols); i++ {
			var res tea.Model
			res, cmd = m.cols[i].Update(msg)
			m.cols[i] = res.(column)
			cmds = append(cmds, cmd)
		}
		m.loaded = true
		return m, tea.Batch(cmds...)
	case Form:
		m.statusMsg = "Task added successfully!"
		m.clearMsgCmd = tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
		return m, tea.Batch(
			m.cols[m.focused].Set(msg.index, msg.CreateTask()), // task addition
			m.clearMsgCmd, // timer command to clear the message
		)
	case clearMessage:
		m.statusMsg = ""
	case moveMsg:
		return m, m.cols[m.focused.getNext()].Set(APPEND, msg.Task)
	case deleteTaskMessage:
		m.statusMsg = "Task deleted successfully!"
		return m, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
	case taskMovedMessage:
		m.statusMsg = "Task moved successfully!"
		return m, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, keys.Left):
			m.cols[m.focused].Blur()
			m.focused = m.focused.getPrev()
			m.cols[m.focused].Focus()
		case key.Matches(msg, keys.Right):
			m.cols[m.focused].Blur()
			m.focused = m.focused.getNext()
			m.cols[m.focused].Focus()
		}
	}
	res, cmd := m.cols[m.focused].Update(msg)
	if _, ok := res.(column); ok {
		m.cols[m.focused] = res.(column)
	} else {
		return res, cmd
	}
	return m, cmd
}

// Changing to pointer receiver to get back to this model after adding a new task via the form...
// Otherwise I would need to pass this model along to the form and it becomes highly coupled to the other models.
func (m *Board) View() string {
	if m.quitting {
		return ""
	}
	if !m.loaded {
		return "loading..."
	}

	statusUI := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")). // green color
		Render(m.statusMsg)

	// calling View() on all the columns
	// this is how each column is rendered inside the board
	board := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.cols[todo].View(),
		m.cols[inProgress].View(),
		m.cols[done].View(),
	)
	return lipgloss.JoinVertical(lipgloss.Left, board, m.help.View(keys), statusUI)
}
