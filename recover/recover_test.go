package recover

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// TestGetRequest tests the web server by creating a web request
func TestGetRequest(t *testing.T) {
	for k, v := range testTable.getRequest {
		req := httptest.NewRequest(http.MethodGet, k, nil)
		w := httptest.NewRecorder()

		switch k {
		case "/":
			homeHandler(w, req)
		case "/panic/":
			panicHandler(w, req)
		default:
			t.Fatal("invalid url")
		}

		resp := w.Result()
		defer resp.Body.Close()

		splits := strings.Split(v, ",")
		expectedStatusCode, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		if expectedStatusCode != resp.StatusCode {
			t.Fatalf("expected %d but got %d. k: %s v: %s", expectedStatusCode, resp.StatusCode, k, v)
		}

		var data bytes.Buffer
		_, err = io.Copy(&data, resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if data.String() != splits[1] {
			t.Fatalf("expected %s but got %s. k: %s v: %s", splits[1], data.String(), k, v)
		}
	}
}

// TestParseTrace tests the parseTrace function
func TestParseTrace(t *testing.T) {
	for k, v := range testTable.parseTrace {
		linkedTrace := parseTrace(k)
		if len(linkedTrace) != len(v) {
			t.Fatalf("expected len %d but got %d. k: %s v: %s", len(v), len(linkedTrace), k, v)
		}

		if linkedTrace != v {
			t.Logf("%s", getUnmatch(linkedTrace, v))
			t.Fatalf("expected %s but got %s. k: %s v: %s", v, linkedTrace, k, v)
		}
	}
}

// helps debug unmatch strings by returning the first string starting from the index of the non-matching character
func getUnmatch(a, b string) string {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i:]
		}
	}
	return ""
}
