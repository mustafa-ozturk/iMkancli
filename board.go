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

func (b *Board) Init() tea.Cmd {
	return nil
}

func (b *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		b.help.Width = msg.Width - margin
		for i := 0; i < len(b.cols); i++ {
			var res tea.Model
			res, cmd = b.cols[i].Update(msg)
			b.cols[i] = res.(column)
			cmds = append(cmds, cmd)
		}
		b.loaded = true
		return b, tea.Batch(cmds...)
	case Form:
		b.statusMsg = "Task added successfully!"
		b.clearMsgCmd = tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
		return b, tea.Batch(
			b.cols[b.focused].Set(msg.index, msg.CreateTask()), // task addition
			b.clearMsgCmd, // timer command to clear the message
		)
	case clearMessage:
		b.statusMsg = ""
	case moveMsg:
		return b, b.cols[b.focused.getNext()].Set(APPEND, msg.Task)
	case deleteTaskMessage:
		b.statusMsg = "Task deleted successfully!"
		return b, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
	case taskMovedMessage:
		b.statusMsg = "Task moved successfully!"
		return b, tea.Tick(2*time.Second, func(_ time.Time) tea.Msg {
			return clearMessage{}
		})
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			b.quitting = true
			b.saveTasks()
			return b, tea.Quit
		case key.Matches(msg, keys.Left):
			b.moveFocus(-1)
		case key.Matches(msg, keys.Right):
			b.moveFocus(1)
		}
	}
	res, cmd := b.cols[b.focused].Update(msg)
	if _, ok := res.(column); ok {
		b.cols[b.focused] = res.(column)
	} else {
		return res, cmd
	}
	return b, cmd
}

// Changing to pointer receiver to get back to this model after adding a new task via the form...
// Otherwise I would need to pass this model along to the form and it becomes highly coupled to the other models.
func (b *Board) View() string {
	if b.quitting {
		return ""
	}
	if !b.loaded {
		return "loading..."
	}

	statusUI := lipgloss.NewStyle().
		Foreground(lipgloss.Color(DefaultTheme.StatusTextColor)).
		Render(b.statusMsg)

	// calling View() on all the columns
	// this is how each column is rendered inside the board
	board := lipgloss.JoinHorizontal(
		lipgloss.Left,
		b.cols[todo].View(),
		b.cols[inProgress].View(),
		b.cols[done].View(),
	)
	return lipgloss.JoinVertical(lipgloss.Left, board, b.help.View(keys), statusUI)
}

func (b *Board) moveFocus(direction int) {
	b.cols[b.focused].Blur()

	if direction < 0 { // left
		b.focused = b.focused.getPrev()
	} else { // right
		b.focused = b.focused.getNext()
	}

	b.syncTaskSelection()

	b.cols[b.focused].Focus()
}

func (b *Board) syncTaskSelection() {
	// ensure only one column's tasks are selected
	for i := range b.cols {
		if i == int(b.focused) {
			// if there are tasks in focused column, select the first one
			if len(b.cols[i].list.Items()) > 0 {
				b.cols[i].list.Select(0)
			}
		} else {
			// unselect everything else
			b.cols[i].list.Select(-1)
		}
	}
}
