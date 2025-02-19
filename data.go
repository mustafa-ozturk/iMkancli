package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	b.updateColumnTitles()

	// focus the correct column
	for i := range b.cols {
		if b.cols[i].status == b.focused {
			b.cols[i].focus = true
			break
		}
	}
}

func (b *Board) updateColumnTitles() {
	b.cols[todo].list.Title = fmt.Sprintf("To Do (%d)", len(b.cols[todo].list.Items()))
	b.cols[inProgress].list.Title = fmt.Sprintf("To Do (%d)", len(b.cols[inProgress].list.Items()))
	b.cols[done].list.Title = fmt.Sprintf("To Do (%d)", len(b.cols[done].list.Items()))
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

	// TODO: Investigate why InsertItem() works while SetItems() caused task loss.
	// Using InsertItem() ensures tasks persist across app restarts.
	for _, task := range tasksData.Tasks {
		b.cols[task.Status].list.InsertItem(len(b.cols[task.Status].list.Items()), Task{
			status:      task.Status,
			title:       task.Title,
			description: task.Description,
		})
	}
}

func (b *Board) saveTasks() {
	// build struct
	todoItems := b.cols[todo].list.Items()
	inProgressItems := b.cols[inProgress].list.Items()
	doneItems := b.cols[done].list.Items()

	allItems := append(append(todoItems, inProgressItems...), doneItems...)

	var tasksData TasksData

	for _, item := range allItems {
		// type assertion allows us to use task as a Task
		if task, ok := item.(Task); ok {
			tasksData.Tasks = append(tasksData.Tasks, TaskJSON{
				Status:      task.status,
				Title:       task.title,
				Description: task.description,
			})
		}
	}

	// encode to json
	jsonData, err := json.Marshal(tasksData)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
		return
	}

	// overwrite json file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
		return
	}

	log.Println("saveTasks() was called")
}
