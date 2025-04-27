package gbld

import (
	"os"
	"path/filepath"
	"strings"
)

func change_ext(path string, new_ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + new_ext
}

func read_ignore_ext(dir string, name string) ([]byte, string, error) {
	entries, e := os.ReadDir(dir)

	if e != nil {
		return nil, "", e
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // skip directories
		}

		if strings.HasPrefix(entry.Name(), name) {
			data, e := os.ReadFile(filepath.Join(dir, entry.Name()))
			return data, entry.Name(), e
		}
	}

	return nil, "", os.ErrNotExist
}
