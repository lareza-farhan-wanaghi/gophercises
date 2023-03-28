package recover

import (
	"fmt"
	"log"
	"net/http"
)

// generalPanicRecovery recovers panic stats of the web server by hijacking the connection
func generalPanicRecovery(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong")
	}
}

// homeHandler handles the root url
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// homeHandler handles the panic url, which triggers panic in the web server
func panicHandler(w http.ResponseWriter, r *http.Request) {
	defer generalPanicRecovery(w, r)
	panic("Oh no!")
}

// RunWebserver runs the web server simulating panic recovery
func RunWebserver() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicHandler)
	mux.HandleFunc("/", homeHandler)

	log.Printf("Stating a web server on port 3000")
	http.ListenAndServe(":3000", mux)
}
