package transform

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// TestPrimitive tests the Primitive function
func TestPrimitive(t *testing.T) {
	for _, v := range testTable.primitive {
		splits := strings.Split(v, ",")
		mode, err := strconv.Atoi(splits[1])
		if err != nil {
			t.Fatal(err)
		}

		n, err := strconv.Atoi(splits[2])
		if err != nil {
			t.Fatal(err)
		}

		in := splits[0]
		extIndex := strings.LastIndex(in, ".")
		out := fmt.Sprintf("%s%s%s", in[:extIndex], "_out", in[extIndex:])

		err = Primitive(in, out, mode, n)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(out); errors.Is(err, os.ErrNotExist) {
			t.Fatal("the out file is not exist")
		}

		err = os.RemoveAll(out)
		if err != nil {
			t.Fatal(err)
		}
	}
}
