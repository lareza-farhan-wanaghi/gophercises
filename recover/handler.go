package recover

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// homeHandler handles the root url
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// homeHandler handles the panic url, which triggers panic in the web server
func panicHandler(w http.ResponseWriter, r *http.Request) {
	defer generalPanicRecovery(w, r)
	panic("Oh no!")
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	defer generalPanicRecovery(w, r)
	if !isDevServer {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/debug")
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, f)
	if err != nil {
		panic(err)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	err = hightlightGoCode(w, buffer.String(), [2]int{id, id})
	if err != nil {
		panic(err)
	}
}
