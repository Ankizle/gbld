package gbld

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func change_ext(path string, new_ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + new_ext
}

func get_ext(name string) string {
	sp := strings.Split(name, ".")[1:]
	return "." + strings.Join(sp, ".")
}

func read_all_ext(dir string, name string, ext_prefs []string) ([]byte, string, error) {
	entries, e := os.ReadDir(dir)

	if e != nil {
		return nil, "", e
	}

	var files_at_exts = make([]string, len(ext_prefs))

	for _, entry := range entries {
		if entry.IsDir() {
			continue // skip directories
		}

		ename := entry.Name()
		if strings.HasPrefix(ename, name) && slices.Contains(ext_prefs, get_ext(ename)) {
			files_at_exts[slices.Index(ext_prefs, get_ext(ename))] = filepath.Join(dir, entry.Name())
		}
	}

	for _, path := range files_at_exts {
		if path == "" {
			continue
		}
		data, e := os.ReadFile(path)
		return data, filepath.Base(path), e
	}

	return nil, "", os.ErrNotExist
}

func reverse[T any](a []T) []T {
	var output = make([]T, len(a))
	for i := 0; i < len(a); i++ {
		output[len(a)-i-1] = a[i]
	}
	return output
}
