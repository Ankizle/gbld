package gbld_c

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func GetSharedPreExt(os string) (prefix string, ext string) {
	switch os {
	case "linux":
		prefix = "lib"
		ext = ".so"
	case "windows":
		ext = ".dll"
	}

	return prefix, ext
}

func Shared(os string, f string) gbld.File {
	prefix, ext := GetSharedPreExt(os)

	shared_path := gbld_fs.ChangeExt(f, ext)
	shared_path = gbld_fs.AddPrefix(shared_path, prefix)
	return gbld.NewFile(shared_path)
}

func SharedVersion(f gbld.File, major uint, minor uint) gbld.File {
	ext := filepath.Ext(f.Path())

	if major != 0 {
		ext += fmt.Sprintf(".%d", major)
	}
	if minor != 0 {
		ext += fmt.Sprintf(".%d", minor)
	}

	return gbld.NewFile(gbld_fs.ChangeExt(f.Path(), ext))
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
