package main

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	task := Task{status: todo, title: "Test Task", description: "This is a test"}

	if task.status != todo {
		t.Errorf("Expected status to be todo, got %v", task.status)
	}
	if task.title != "Test Task" {
		t.Errorf("Expected title to be 'Test Task', got %v", task.title)
	}
	if task.description != "This is a test" {
		t.Errorf("Expected description to be 'This is a test', got %v", task.description)
	}
}

func TestTaskMoveNext(t *testing.T) {
	task := Task{status: todo}

	task.status = task.status.getNext()

	if task.status != inProgress {
		t.Errorf("Expected status to be inProgress, got %v", task.status)
	}

	task.status = task.status.getNext()

	if task.status != done {
		t.Errorf("Expected status to be done, got %v", task.status)
	}
}
