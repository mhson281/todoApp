/*
Copyright © 2024 Minh Son <son.minh@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todoApp",
	Short: "An application to manage your todo list",
	Long: `The application will help you Add, List, Update and Delete your todo list.
	Example:
	$tasks add "Buy milk"
	$ tasks list
ID    Task                                                Created
1     Tidy up my desk                                     a minute ago
3     Change my keyboard mapping to use escape/control    a few seconds ago
`,
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
}
