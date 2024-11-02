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

	err = writer.Write([]string{fmt.Sprintf("%d. %s", taskID, task)})
	if err != nil {
		return err
	}

	fmt.Printf("Task %s with ID #%d has been added to todo list", task, taskID)
	return nil
}

func ReadTaskFromCSV() ([]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []string
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		tasks = append(tasks, record...)
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
	table.SetHeader([]string{"ID", "Task"})

	for _, task := range tasks {
		id, content := splitTask(task)
		table.Append([]string{id, content})
	}

	table.Render()
	return nil
}

func splitTask(task string) (string, string){
	parts := strings.Split(task, ".")
	id := parts[0]
	content := parts[1]

	return id, content
}

