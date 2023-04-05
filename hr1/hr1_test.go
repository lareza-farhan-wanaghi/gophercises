package hr1

import (
	"strconv"
	"strings"
	"testing"
)

// TestCamelcase tests the camelcase function
func TestCamelcase(t *testing.T) {
	for k, v := range testTable.camelcase {
		result := Camelcase(k)

		if result != v {
			t.Fatalf("Expected %d but got %d. data being inspected: '%s'",
				v, result, k)
		}
	}
}

// TestCaesarChiper tests the caesarChiper function
func TestCaesarChiper(t *testing.T) {
	for k, v := range testTable.caesarChiper {
		split := strings.Split(k, ",")
		offset, err := strconv.ParseInt(split[1], 10, 32)
		if err != nil {
			t.Fatal(err)
		}

		result := CaesarChiper(split[0], int32(offset))
		if result != v {
			t.Fatalf("Expected '%s' but got '%s'. data being inspected: '%s'",
				v, result, k)
		}
	}
}
