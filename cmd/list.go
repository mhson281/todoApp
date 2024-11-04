package cmd

import (
  "fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)
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
