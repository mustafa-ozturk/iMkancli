package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

func (s status) getNext() status {
	if s == done {
		return todo
	}
	return s + 1
}

func (s status) getPrev() status {
	if s == todo {
		return done
	}
	return s - 1
}

const margin = 4

var board *Board

const (
	todo status = iota
	inProgress
	done
)

type clearMessage struct{}
type taskMovedMessage struct{}

func NewBoard() *Board {
	help := help.New()
	help.ShowAll = true

	board := &Board{
		help:    help,
		focused: todo,
	}

	board.initLists()
	board.syncTaskSelection()

	return board
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	board = NewBoard()
	p := tea.NewProgram(board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
