package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	quitehn "github.com/lareza-farhan-wanaghi/gophercises/quiet_hn"
)

// main provides the entry point of the app
func main() {
	client := quitehn.NewClient(30, 10*time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stories := client.GetTopStories(30)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stories)
	})

	fmt.Println("starting a server at port 80")
	http.ListenAndServe(":80", nil)
}
