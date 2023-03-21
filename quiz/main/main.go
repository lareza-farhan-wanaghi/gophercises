package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// main provides the entry point of the app
func main() {
	timeFlag := flag.Int("t", 30, "Specifies time of the quiz")
	shuffleFlag := flag.Bool("s", false, "Specifies whether the quiz need to be shuffled")
	flag.Parse()

	data, err := getQuenstions("problems.csv", shuffleFlag)
	if err != nil {
		panic(err)
	}
	simulateQuiz(data, timeFlag)
}

// simulateQuiz simulates the desired quiz, which will read user inputs and respond to them
func simulateQuiz(data [][]string, timeFlag *int) {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("You will be given %d seconds to answer the following questions.\nPress ENTER to start the quiz!\n", *timeFlag)
	scanner.Scan()

	timer := time.NewTimer(time.Duration(*timeFlag) * time.Second)
	isTimesup := false
	go func() {
		<-timer.C
		fmt.Printf("You have correctly answered %d over %d questions!\n", score, len(data))
		os.Exit(0)
	}()

	for i := 0; i < len(data) && !isTimesup; i++ {
		fmt.Printf("What %s, sir?\n", data[i][0])
		scanner.Scan()
		answer := scanner.Text()
		if strings.TrimSpace(answer) == data[i][1] {
			score += 1
		}
	}
	fmt.Printf("You have correctly answered %d over %d questions!\n", score, len(data))
}

// getQuenstions returns question and answer pairs from a given CSV file
func getQuenstions(filename string, shuffleFlag *bool) ([][]string, error) {
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

	if *shuffleFlag {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	}
	return data, nil
}
