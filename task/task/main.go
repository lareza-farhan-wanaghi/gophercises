package main

import "github.com/lareza-farhan-wanaghi/gophercises/task"

func main() {
	err := task.Execute()
	if err != nil {
		panic(err)
	}
}
