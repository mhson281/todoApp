/*
Copyright © 2024 Minh Son <son.minh@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todoApp",
	Short: "An application to manage your todo list",
	Long: `This application will help you Add, List, Update and Delete your todo list.
	Example:
	$tasks add "Buy milk"
	$ tasks list
	ID    Task                                                Created
	1     Tidy up my desk                                     a minute ago
	3     Change my keyboard mapping to use escape/control    a few seconds ago
`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the task to be added")
			os.Exit(1)
		}
		task := args[0]
		if err := AddTaskToCSV(task); err != nil {
			fmt.Printf("Error adding task: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task has been added: %s\n", task)
	},
}

var listCmd = &cobra.Command{
  Use:	 "list",
  Short: "list all tasks in to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		if err := PrintTaskTable(); err != nil {
			fmt.Printf("Error displaying tasks: %v\n", err)
		}
	},
}

var markCompleteCmd = &cobra.Command {
	Use: "complete",
	Short: "mark task with provided ID as complete",
	Run: func (cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the task id to be marked completed")
		}
		taskId := args[0]
		if err := MarkTaskComplete(taskId); err != nil {
			fmt.Printf("Unable to mark task ID #%s as completed\n", taskId)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todoApp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(markCompleteCmd)
}
