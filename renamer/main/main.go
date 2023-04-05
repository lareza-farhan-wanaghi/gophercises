package main

import (
	"os"
	"path/filepath"

	"github.com/lareza-farhan-wanaghi/gophercises/renamer"
)

// main provides the entry point of the app
func main() {
	if len(os.Args) < 3 {
		panic("required arguments: [pattern] [to] [rootPath]")
	}
	pattern := os.Args[1]
	to := os.Args[2]
	rootPath := os.Args[3]

	err := filepath.Walk(rootPath, renamer.RenameFiles(&pattern, &to))

	if err != nil {
		panic(err)
	}
}
