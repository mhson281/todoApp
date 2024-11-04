package cmd

import (
	"fmt"
	"os"
)


func ClearToDoList() error {
	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()


	fmt.Println("All tasks have been cleared from the to-do list")
	return nil
}
