package gbld_c

import (
	"path/filepath"
	"strings"

	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func Shared(os string, f string) gbld.File {
	src_path := f
	var prefix string
	var ext string

	switch os {
	case "linux":
		prefix = "lib"
		ext = ".so"
	case "windows":
		ext = ".dll"
	}

	shared_path := gbld_fs.ChangeExt(src_path, ext)
	shared_path = gbld_fs.AddPrefix(shared_path, prefix)
	return gbld.NewFile(shared_path)
}

func FromSharedFile(os string, f gbld.File) string {
	src_path := filepath.Base(f.Path())
	var prefix string

	switch os {
	case "linux":
		prefix = "lib"
	}

	libname := strings.TrimSuffix(gbld_fs.ChangeExt(src_path, ""), ".")
	libname = strings.TrimPrefix(libname, prefix)
	return libname
}
