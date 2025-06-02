package gbld_c

import (
	"path/filepath"
	"strings"
)

func change_extension(name string, new_ext string) string {
	ext := filepath.Ext(name)
	new_name := strings.TrimSuffix(name, ext)
	return new_name + new_ext
}
