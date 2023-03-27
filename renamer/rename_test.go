package renamer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRenameFile tests the renameFile function
func TestRenameFile(t *testing.T) {
	for k, v := range testTable.renameFile {
		err := os.RemoveAll(testDir)
		if err != nil {
			t.Fatal(err)
		}

		err = os.MkdirAll(testDir, 0777)
		if err != nil {
			t.Fatal(err)
		}

		splits := strings.Split(k, ",")
		for _, s := range splits[:len(splits)-2] {
			fn := filepath.Join(testDir, s)
			err = os.MkdirAll(filepath.Dir(fn), 0777)
			if err != nil {
				t.Fatal(err)
			}

			_, err = os.Create(fn)
			if err != nil {
				t.Fatal(err)
			}
		}

		pattern := splits[len(splits)-2]
		to := splits[len(splits)-1]
		err = filepath.Walk(testDir, RenameFiles(&pattern, &to))
		if err != nil {
			t.Fatal(err)
		}

		splits, i := strings.Split(v, ","), 0
		err = filepath.Walk(testDir, func(path string, info fs.FileInfo, err error) error {
			if info.Mode().IsDir() {
				return nil
			}

			fn := filepath.Base(path)
			if fn != splits[i] {
				t.Fatalf("expected %s but got %s. k: %s v: %s", splits[i], fn, k, v)
			}
			i += 1
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}

		if i != len(splits) {
			t.Fatalf("expected inspecting %d files but got %d. k: %s v: %s", len(splits), i, k, v)
		}
	}
}
