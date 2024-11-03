package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const csvFilePath = "./tasks.csv"

func AddTaskToCSV(task string) error {
	tasks, err := ReadTaskFromCSV()
	if err != nil && os.IsExist(err) {
		return err
	}

	// Determine the next task number
	taskID := len(tasks) + 1

	file, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"[ ]", fmt.Sprintf("%d", taskID), task})
	if err != nil {
		return err
	}

	fmt.Printf("Task %s with ID #%d has been added to todo list", task, taskID)
	return nil
}

func ReadTaskFromCSV() ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks [][]string
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		tasks = append(tasks, record)
	}

	return tasks, nil
}

func PrintTaskTable() error {
	tasks, err := ReadTaskFromCSV()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status","ID", "Task"})

	for _, task := range tasks {
		if len(task) == 3 {
			table.Append(task)
		}
	}

	table.Render()
	return nil
}

func splitTaskWithStatus(task string) (string, string, string){
	parts := strings.SplitN(task, ",", 3)

	if len(parts) < 2 {
		return "[ ]", parts[0], ""
	}

	status := parts[0]
	id := parts[1]
	content := parts[2]

	return status, id, content
}

func MarkTaskComplete(taskID string) error {
	tasks, err := ReadTaskFromCSV()
	if err != nil {
		return err
	}

	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	completedTask := false

	for _, task := range tasks {
		status, id, content := task[0], task[1], task[2]
		if id == taskID {
			status = "[x]"
			completedTask = true
		}

		err := writer.Write([]string{status, id, content})
		if err != nil {
			return err
		}

	}

	if completedTask {
		fmt.Printf("Taskd ID #%s has been marked as complete.\n", taskID)
	} else {
		fmt.Printf("Taskd ID #%s was not found in the to-do list.\n", taskID)
	}

	return nil
}

func RemoveTask(taskID string) error {
	tasks, err := ReadTaskFromCSV()
	if err != nil {
		return err
	}

	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	removedTask := false

	for _, task := range tasks {
		status, id, content := task[0], task[1], task[2]
		if id == taskID {
			removedTask = true
			continue
		}

		err := writer.Write([]string{status, id, content})
		if err != nil {
			return err
		}
	}

	if  removedTask {
		fmt.Printf("Task ID #%s has been removed.\n", taskID)
	} 

	return nil
}



