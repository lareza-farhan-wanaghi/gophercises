package main

import "github.com/lareza-farhan-wanaghi/gophercises/task"

// main provides the entry point of the app
func main() {
	err := task.Execute()
	if err != nil {
		panic(err)
	}
}
