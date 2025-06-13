package gbld_c

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func Executable(os string, f string) gbld.File {
	src_path := f
	var ext string

	switch os {
	case "linux":
		ext = ""
	case "windows":
		ext = ".exe"
	}

	shared_path := gbld_fs.ChangeExt(src_path, ext)
	return gbld.NewFile(shared_path)
}
