package main

// main provides the entry point of the app
func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
