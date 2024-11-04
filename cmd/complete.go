package cmd

import (
	"fmt"
	"os"
	"encoding/csv"
)

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
