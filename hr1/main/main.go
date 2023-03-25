package main

import (
	"flag"
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/hr1"
)

// main provides the entry point of the app
func main() {
	textFlag := flag.String("t", "camelCase", "Specifies the text that will be inspected")
	offsiteFlag := flag.Int("o", 2, "Specifies the offsite that will be used in the caesar chipper encryption")
	flag.Parse()

	fmt.Printf("inputed text: '%s', CamelCase output: %d, CaesarChiper output: '%s'\n",
		*textFlag, hr1.Camelcase(*textFlag), hr1.CaesarChiper(*textFlag, int32(*offsiteFlag)))

}
