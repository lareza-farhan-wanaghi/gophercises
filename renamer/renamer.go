package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// RenameFiles renames files in the directory recursively that match the pattern
func RenameFiles(pattern, to *string) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil || info.Mode().IsDir() {
			return err
		}

		dn := filepath.Dir(path)
		fn := filepath.Base(path)

		renamedFn, isSuccess := regexReplaceAll(&fn, pattern, to)
		if !isSuccess {
			renamedFn = wordsReplaceAll(&fn, pattern, to)
		}

		if fn == renamedFn {
			return nil
		}

		newPath := filepath.Join(dn, renamedFn)
		if _, err = os.Stat(newPath); !os.IsNotExist(err) {
			fmt.Printf("skip renaming %s to %s since a file is exist in that path\n", path, newPath)
			return err
		}

		err = os.Rename(path, newPath)
		if err != nil {
			return err
		}

		return nil
	}
}

// regexReplaceAll finds words that match to the regex pattern and replaces them with the new words
func regexReplaceAll(target, pattern, to *string) (string, bool) {
	re, err := regexp.Compile(*pattern)
	if err != nil {
		return "", false
	}

	return re.ReplaceAllString(*target, *to), true
}

// wordsReplaceAll finds words that are equal to the pattern and replaces them with the new words
func wordsReplaceAll(target, pattern, to *string) string {
	return strings.ReplaceAll(*target, *pattern, *to)
}
