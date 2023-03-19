package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// main provides the entry point of the app
func main() {
	data, err := getQuenstions("problems.csv")
	if err != nil {
		panic(err)
	}
	simulateQuiz(data)
}

// simulateQuiz simulates the desired quiz, which will read user inputs and respond to them
func simulateQuiz(data [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	for i := 0; i < len(data); i++ {
		fmt.Printf("What %s, sir?\n", (data)[i][0])
		scanner.Scan()
		answer := scanner.Text()
		if strings.TrimSpace(answer) == (data)[i][1] {
			score += 1
		}
	}
	fmt.Printf("You have correctly answered %d over %d questions!\n", score, len(data))
}

// getQuenstions returns question and answer pairs from a given CSV file
func getQuenstions(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
