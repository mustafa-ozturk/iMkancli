package main

import (
	"testing"

	"github.com/charmbracelet/bubbles/list"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	if len(board.cols) != 3 {
		t.Fatalf("Expected board to have 3 columns, got %d", len(board.cols))
	}

	if board.focused != todo {
		t.Errorf("Expected focused column to be 'todo', got %v", board.focused)
	}
}

func TestMoveFocus(t *testing.T) {
	board := NewBoard()

	if board.focused != todo {
		t.Fatalf("Expected to start focused on 'To Do', got %v", board.focused)
	}

	board.moveFocus(1)
	if board.focused != inProgress {
		t.Errorf("Expected focus to move to 'In Progress', got %v", board.focused)
	}

	board.moveFocus(1)
	if board.focused != done {
		t.Errorf("Expected focus to move to 'Done', got %v", board.focused)
	}

	board.moveFocus(-1)
	if board.focused != inProgress {
		t.Errorf("Expected focus to move back to 'In Progress', got %v", board.focused)
	}

	board.moveFocus(-1)
	if board.focused != todo {
		t.Errorf("Expected focus to move back to 'To Do', got %v", board.focused)
	}
}

func TestBoardAddTask(t *testing.T) {
	board := NewBoard()
	task := Task{status: todo, title: "Test Task", description: "Testing"}

	board.cols[todo].list.SetItems([]list.Item{task})

	items := board.cols[todo].list.Items()
	if len(items) != 1 {
		t.Fatalf("Expected 1 task in To Do column, got %d", len(items))
	}

	addedTask, ok := items[0].(Task)
	if !ok {
		t.Fatal("Failed to retrieve task from list")
	}
	if addedTask.Title() != "Test Task" {
		t.Errorf("Expected task title to be 'Test Task', got '%s'", addedTask.Title())
	}
	if addedTask.Description() != "Testing" {
		t.Errorf("Expected task description to be 'Testing', got '%s'", addedTask.Description())
	}
}
