package gbld_fs

import (
	"path/filepath"
	"strings"
)

func ChangeExt(name string, new_ext string) string {
	ext := filepath.Ext(name)
	new_name := strings.TrimSuffix(name, ext)
	return new_name + new_ext
}
