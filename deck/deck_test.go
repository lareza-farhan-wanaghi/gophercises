package deck

import (
	"bufio"
	"os"
	"testing"
)

// TestNewSuitDeck tests the NewSuitDeck function
func TestNewSuitDeck(t *testing.T) {
	for k, v := range testTable.newSuitDeck {
		func() {
			file, err := os.Open(k)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			sd := NewSuitDeck(v...)
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for i := 0; scanner.Scan(); i++ {
				expectedVal := scanner.Text()
				returnVal := sd[i].String()

				if returnVal != expectedVal {
					t.Fatalf("expected '%s' but got '%s'", expectedVal, returnVal)
				}
			}

			if err := scanner.Err(); err != nil {
				t.Fatal(err)
			}
		}()
	}
}
