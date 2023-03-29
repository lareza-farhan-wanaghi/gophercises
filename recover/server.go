package recover

import (
	"flag"
	"log"
	"net/http"
)

var isDevServer bool

// RunServer runs the web server simulating panic recovery
func RunServer() {
	devFlag := flag.Bool("dev", false, "specifies that the server is in the development mode")
	flag.Parse()
	isDevServer = *devFlag

	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/debug/", sourceCodeHandler)

	log.Printf("Stating a web server on port 3000")
	http.ListenAndServe(":3000", mux)
}
