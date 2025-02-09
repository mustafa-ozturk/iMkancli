package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

// Provides the mock data to fill the kanban board

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}

	/*TODO: get these from the json file */
	// Init To Do
	b.cols[todo].list.Title = "To Do"
	// Init in progress
	b.cols[inProgress].list.Title = "In Progress"
	// Init done
	b.cols[done].list.Title = "Done"

	b.loadTasks()

	// focus the correct column
	for i := range b.cols {
		if b.cols[i].status == b.focused {
			b.cols[i].focus = true
			break
		}
	}
}

type TaskJSON struct {
	Status      status `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TasksData struct {
	Tasks []TaskJSON `json:"tasks"`
}

func (b *Board) loadTasks() {
	data, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatalf("Error: Could not read 'data.json': %v", err)
	}

	var tasksData TasksData
	err = json.Unmarshal(data, &tasksData)
	if err != nil {
		log.Fatalf("Error: Failed to parse JSON from 'data.json': %v", err)
	}

	for _, task := range tasksData.Tasks {
		b.cols[task.Status].list.SetItems([]list.Item{
			Task{status: task.Status, title: task.Title, description: task.Description},
		})
	}
}
