package cmd

import (
	"fmt"
	"encoding/csv"
	"os"
)

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
