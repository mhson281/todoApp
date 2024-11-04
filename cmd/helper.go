package cmd

import (
	"encoding/csv"
	"os"
	"strings"

)

const csvFilePath = "./tasks.csv"


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





