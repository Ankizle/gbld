package gbld_fs

import "path/filepath"

func AddPrefix(path string, prefix string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	new_base := prefix + base
	return filepath.Join(dir, new_base)
}
