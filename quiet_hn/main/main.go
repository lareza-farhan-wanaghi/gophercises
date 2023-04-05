package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	quitehn "github.com/lareza-farhan-wanaghi/gophercises/quiet_hn"
)

// main provides the entry point of the app
func main() {
	refreshCacheFlag := flag.Int("r", 10, "Specifies the time, in second, to refresh the story cache")

	defaultNumOfData := 30
	client := quitehn.NewClient(defaultNumOfData, time.Duration(*refreshCacheFlag)*time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nQuery := r.URL.Query().Get("n")
		n, err := strconv.Atoi(nQuery)
		if err != nil {
			n = defaultNumOfData
		}
		log.Println(n)
		stories := client.GetTopStories(n)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stories)
	})

	fmt.Println("starting a server at port 8080")
	http.ListenAndServe(":8080", nil)
}
