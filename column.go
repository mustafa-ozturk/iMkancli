package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const APPEND = -1

type column struct {
	// used to highlight the currently selected column
	focus  bool
	status status
	list   list.Model
	height int
	width  int
}

func (c *column) Focus() {
	c.focus = true
}

func (c *column) Blur() {
	c.focus = false
}

func (c *column) Focused() bool {
	return c.focus
}

func newColumn(status status) column {
	defaultDelegate := list.NewDefaultDelegate()

	defaultDelegate.Styles.SelectedTitle.
		BorderForeground(DefaultTheme.SelectedBorderColor).
		Foreground(DefaultTheme.SelectedTitleColor).
		Bold(true)
	defaultDelegate.Styles.SelectedDesc.
		BorderForeground(DefaultTheme.SelectedBorderColor).
		Foreground(DefaultTheme.SelectedDescColor)

	defaultList := list.New([]list.Item{}, defaultDelegate, 0, 0)
	defaultList.Styles.Title = lipgloss.NewStyle().
		Foreground(DefaultTheme.ColumnTitleColor).
		Background(DefaultTheme.ColumnTitleBgColor).
		Padding(0, 1)

	defaultList.SetShowHelp(false)
	return column{focus: false, status: status, list: defaultList}
}

// Init does initial setup for the column.
func (c column) Init() tea.Cmd {
	return nil
}

// Update handles all the I/O for columns.
func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.setSize(msg.Width, msg.Height)
		c.list.SetSize(msg.Width/margin, msg.Height/2)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Edit):
			if len(c.list.VisibleItems()) != 0 {
				task := c.list.SelectedItem().(Task)
				f := NewForm(task.title, task.description)
				f.index = c.list.Index()
				f.col = c
				return f.Update(nil)
			}
		case key.Matches(msg, keys.New):
			f := newDefaultForm()
			f.index = APPEND
			f.col = c
			return f.Update(nil)
		case key.Matches(msg, keys.Delete):
			msg, cmd := c.DeleteCurrent()
			return c, tea.Batch(
				cmd,                           // ensure list updates
				func() tea.Msg { return msg }, // Convert msg into a cmd
			)
		case key.Matches(msg, keys.Enter):
			msg, cmd := c.MoveToNext()
			return c, tea.Batch(
				cmd,
				func() tea.Msg { return msg },
			)
		}
	}
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c column) View() string {
	// first gets the style for the column then renders it
	return c.getStyle().Render(c.list.View())
}

type deleteTaskMessage struct{}

func (c *column) DeleteCurrent() (tea.Msg, tea.Cmd) {
	if len(c.list.VisibleItems()) > 0 {
		c.list.RemoveItem(c.list.Index())
		var cmd tea.Cmd
		c.list, cmd = c.list.Update(nil)
		return deleteTaskMessage{}, cmd // Notify Board.Update()
	}

	return nil, nil // no deletion happened
}

func (c *column) Set(i int, t Task) tea.Cmd {
	// if index is within bounds, overwrite an existing task
	if i != APPEND {
		return c.list.SetItem(i, t)
	}
	// otherwise append a new task
	return c.list.InsertItem(APPEND, t)
}

func (c *column) setSize(width, height int) {
	c.width = width / margin
}

func (c *column) getStyle() lipgloss.Style {
	if c.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(DefaultTheme.BorderColor)).
			Height(c.height).
			Width(c.width)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Height(c.height).
		Width(c.width)
}

type moveMsg struct {
	Task
}

func (c *column) MoveToNext() (tea.Msg, tea.Cmd) {
	var task Task
	var ok bool
	// If nothing is selected, the SelectedItem will return Nil.
	if task, ok = c.list.SelectedItem().(Task); !ok {
		return nil, nil
	}

	// remove item if done
	if c.status == done {
		c.list.RemoveItem(c.list.Index())
		return deleteTaskMessage{}, nil
	}

	// move item
	c.list.RemoveItem(c.list.Index())
	task.status = c.status.getNext()

	return taskMovedMessage{}, func() tea.Msg { return moveMsg{task} }
}
