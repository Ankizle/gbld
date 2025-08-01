package gbld_fs

import (
	"strings"
)

func ChangeExt(name string, new_ext string) string {
	base_noext := strings.Split(name, ".")[0]
	return base_noext + new_ext
}
