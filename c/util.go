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

func add_prefix(path string, prefix string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	new_base := prefix + base
	return filepath.Join(dir, new_base)
}
