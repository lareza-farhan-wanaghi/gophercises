package main

import (
	"flag"

	"github.com/lareza-farhan-wanaghi/gophercises/quiz"
)

// main provides the entry point of the app
func main() {
	timeFlag := flag.Int("t", 30, "Specifies time of the quiz")
	shuffleFlag := flag.Bool("s", false, "Specifies whether the quiz need to be shuffled")
	flag.Parse()

	quizSimulator, err := quiz.NewQuizSimulator("problems.csv", *shuffleFlag, *timeFlag)
	if err != nil {
		panic(err)
	}
	quizSimulator.SimulateQuiz()
}
