
package cmd

import (
	"os"
	"fmt"
	"encoding/csv"
) 
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
