package secret

import (
	"strings"
	"testing"
)

// TestGet tests the Get function of the fileFault struct
func TestGet(t *testing.T) {
	for k, v := range testTable.get {
		splits := strings.Split(k, ",")

		ff := FileFault(testEncodingKey, testFilePath)
		for _, sp := range splits {
			kvSplits := strings.Split(sp, ":")
			err := ff.Set(kvSplits[0], kvSplits[1])
			if err != nil {
				t.Fatal(err)
			}
		}

		splits = strings.Split(v, ":")
		data, err := ff.Get(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		if data != splits[1] {
			t.Fatalf("expected %s but got %s. k:%s v:%s", splits[1], data, k, v)
		}

		ff.empty()
	}
}

// TestGetAll tests the GetAll function of the fileFault struct
func TestGetAll(t *testing.T) {
	for k, v := range testTable.getAll {
		splits := strings.Split(k, ",")

		ff := FileFault(testEncodingKey, testFilePath)
		for _, sp := range splits {
			kvSplits := strings.Split(sp, ":")
			err := ff.Set(kvSplits[0], kvSplits[1])
			if err != nil {
				t.Fatal(err)
			}
		}

		splits = strings.Split(v, ",")
		data, err := ff.GetAll()
		if err != nil {
			t.Fatal(err)
		}

		for _, sp := range splits {
			kvSplits := strings.Split(sp, ":")
			val := data[kvSplits[0]]
			if val != kvSplits[1] {
				t.Fatalf("expected %s but got %s. k:%s v:%s", kvSplits[1], val, k, v)
			}
		}

		ff.empty()
	}
}
