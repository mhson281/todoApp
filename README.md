# To-Do List CLI App

A simple command-line application to manage a to-do list. This CLI app allows you to add, list, mark tasks as complete, remove tasks, and clear the entire to-do list, all stored in a CSV file.

## Features

- **Add a Task**: Add a new task to the to-do list.
- **List Tasks**: Display all tasks in a formatted table.
- **Mark Task as Complete**: Mark a specific task as complete.
- **Remove a Task**: Remove a specific task by its ID.
- **Clear To-Do List**: Clear all tasks from the list.

## Requirements

- [Go](https://golang.org/) 1.16+
- [Cobra](https://github.com/spf13/cobra)
- [tablewriter](https://github.com/olekukonko/tablewriter)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/todo-cli-app.git
    cd todo-cli-app
    ```

2. Install dependencies:
    ```bash
    go get -u github.com/spf13/cobra
    go get -u github.com/olekukonko/tablewriter
    ```

3. Build the app:
    ```bash
    go build -o todoApp main.go
    ```

4. Run the app:
    ```bash
    ./todoApp
    ```

## Usage

### 1. Add a Task

Add a new task to the to-do list. Each task will be assigned a unique ID.

**Example:**

```bash
./todoApp add "Buy groceries"
```
### 2. List All Tasks

Display all tasks in a formatted table, showing the task status, ID, and description.

```bash
./todoApp list
```

Example Output:

```css

+--------+----+---------------------+
| Status | ID | Task                |
+--------+----+---------------------+
| [ ]    | 1  | Buy groceries       |
| [x]    | 2  | Walk the dog        |
| [ ]    | 3  | Finish project      |
+--------+----+---------------------+
```

### 3. Mark Task as Complete

Mark a task as complete by specifying its ID. The status will change from [ ] to [x].


```bash
./todoApp complete <taskID>
```

Example:

```
./todoApp complete 5

```bash


Output:


```bash
Taskd ID #5 has been marked as complete.
```


### 4. Remove a Task

Remove a task from the to-do list by its ID.

```
./todoApp remove <taskID>

```bash


Example:

```
./todoApp remove 5

```bash


Output:

```
Task ID #5 has been removed.
```

### 5. Clear All Tasks

Clear all tasks from the to-do list, leaving it empty.

```
./todoApp clear
```bash


Output:

```
All tasks have been cleared from the to-do list.

```

### Code Overview

    AddTaskToCSV: Adds a new task to tasks.csv with a unique ID and an initial status of [ ].
    PrintTaskTable: Reads all tasks from tasks.csv and displays them in a formatted table using tablewriter.
    MarkTaskComplete: Marks a task as complete by updating the status to [x].
    RemoveTaskByID: Removes a task by its ID from tasks.csv.
    ClearToDoList: Clears all tasks by truncating tasks.csv.

File Structure

```
.
├── main.go           # Entry point for the CLI app
├── cmd               # Directory for Cobra commands
│   ├── add.go        # Code for the "add" command
│   ├── list.go       # Code for the "list" command
│   ├── complete.go   # Code for the "complete" command
│   ├── remove.go     # Code for the "remove" command
│   └── clear.go      # Code for the "clear" command
└── tasks.csv         # CSV file where tasks are stored

```plaintext


Contributing

Feel free to fork the repository and submit pull requests for new features, bug fixes, or improvements.
License

This project is licensed under the MIT License.
