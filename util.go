package gbld

import (
	"path/filepath"
	"strings"
)

func change_ext(path string, new_ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + new_ext
}
