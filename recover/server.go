package recover

import (
	"flag"
	"log"
	"net/http"
)

var isDevServer bool

// RunServer runs the web server simulating panic recovery
func RunServer() {
	devFlag := flag.Bool("dev", true, "specifies that the server is in the development mode")
	flag.Parse()
	isDevServer = *devFlag

	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/debug/", sourceCodeHandler)

	log.Printf("Stating a web server on port 8080")
	http.ListenAndServe(":8080", mux)
}
