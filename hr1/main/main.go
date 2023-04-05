package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lareza-farhan-wanaghi/gophercises/hr1"
)

// main provides the entry point of the app
func main() {
	textInput := os.Args[1]
	offsiteInput, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Input text: %s \nWord count: %d \nCaesar chiper result: %s\n",
		textInput, hr1.Camelcase(textInput), hr1.CaesarChiper(textInput, int32(offsiteInput)))

}
