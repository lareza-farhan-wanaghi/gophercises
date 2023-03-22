package task

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// init adds other commands to the root command. It is called by the cobra library before any other command functions
func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completedCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to the task list",
	Run: func(cob *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := taskDb.addTask(task)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Added \"%s\" to your task list!\n", task)
	},
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a simple and fast task manager",
	Long: `A simple and fast task manager to help you manage your tasks. 
	It uses boldDB database in the current directory to store the data`,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Run: func(cob *cobra.Command, args []string) {
		tasks, err := taskDb.listTask()
		if err != nil {
			panic(err)
		}

		if len(tasks) == 0 {
			println("You currently have no tasks!")
			return
		}

		println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Title)
		}
	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark the task as completed by its ID",
	Run: func(cob *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		tm, err := taskDb.doTask(id - 1)
		if err != nil {
			panic(err)
		}

		if tm == nil {
			fmt.Printf("Sorry, we can't find the task with id %d you specified.\n", id)
			return
		}

		fmt.Printf("You have completed the \"%s\" task!\n", tm.Title)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove the task by its ID",
	Run: func(cob *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		tm, err := taskDb.deleteTask(id - 1)
		if err != nil {
			panic(err)
		}

		if tm == nil {
			fmt.Printf("Sorry, we can't find the task with id %d you specified.\n", id)
			return
		}
		fmt.Printf("You have deleted the \"%s\" task!\n", tm.Title)
	},
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List today's completed tasks",
	Run: func(cob *cobra.Command, args []string) {
		completedTasks, err := taskDb.completedTask()
		if err != nil {
			panic(err)
		}
		if len(completedTasks) == 0 {
			println("You don't have any completed tasks today!")
			return
		}

		fmt.Println("You have finished the following tasks today:")
		for i, task := range completedTasks {
			fmt.Printf("%d. %s\n", i+1, task.Title)
		}
	},
}
