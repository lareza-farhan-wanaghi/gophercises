package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type quizSimulator struct {
	data     [][]string
	time     int
	shuffled bool
}

// simulateQuiz simulates the desired quiz, which will read user inputs and respond to them
func (q *quizSimulator) SimulateQuiz() {
	if q.shuffled {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(q.data), q.shuffleData)
	}

	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("You will be given %d seconds to answer the following questions.\nPress ENTER to start the quiz!\n", q.time)
	scanner.Scan()

	timer := time.NewTimer(time.Duration(q.time) * time.Second)
	isTimesup := false
	go func() {
		<-timer.C
		fmt.Printf("You have correctly answered %d over %d questions!\n", score, len(q.data))
		os.Exit(0)
	}()

	for i := 0; i < len(q.data) && !isTimesup; i++ {
		fmt.Printf("What %s, sir?\n", q.data[i][0])
		scanner.Scan()
		answer := scanner.Text()
		if strings.TrimSpace(answer) == q.data[i][1] {
			score += 1
		}
	}
	fmt.Printf("You have correctly answered %d over %d questions!\n", score, len(q.data))
}

// shuffleData shuffles the quiz data
func (q *quizSimulator) shuffleData(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

// getQuenstions returns question and answer pairs from a given CSV file
func NewQuizSimulator(filename string, shuffled bool, time int) (*quizSimulator, error) {
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

	result := &quizSimulator{
		data:     data,
		shuffled: shuffled,
		time:     time,
	}
	return result, nil
}
