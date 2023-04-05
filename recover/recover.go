package recover

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"runtime/debug"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// generalPanicRecovery recovers panic stats of the web server by hijacking the connection
func generalPanicRecovery(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if isDevServer {
			trace := string(debug.Stack())
			linkedTrace := fmt.Sprintf("<body><p style='white-space:pre'>%s\n\n%s</p></body>", err, parseTrace(trace))
			fmt.Fprintln(w, linkedTrace)
		} else {
			fmt.Fprintf(w, "Something went wrong")
		}
	}
}

// parseTrace finds raw source code links and turns them into tags linking to the source codes
func parseTrace(trace string) string {
	re := regexp.MustCompile(`((\/[^\s]+):([0-9]+).*)`)
	result := re.ReplaceAllString(trace, "<a href='/debug${2}?id=${3}'>${1}</a>")

	return result
}

// hightlightGoCode formats the string with the chroma syntax highlighting and writes it to the writer
func hightlightGoCode(w io.Writer, code string, highlightLines ...[2]int) error {
	lexer := lexers.Get("go")
	if lexer == nil {
		return fmt.Errorf("error while highlighting")
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("base16-snazzy")
	if style == nil {
		return fmt.Errorf("error while highlighting")
	}

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return err
	}

	formatter := html.New(html.Standalone(true), html.WithClasses(true), html.HighlightLines(highlightLines))
	err = formatter.Format(w, style, iterator)

	return err
}
