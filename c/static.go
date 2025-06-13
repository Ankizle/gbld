package gbld_c

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func Static(os string, f string) gbld.File {
	src_path := f
	var prefix string
	var ext string

	switch os {
	case "linux":
		prefix = "lib"
		ext = ".a"
	case "windows":
		ext = ".lib"
	}

	static_path := gbld_fs.ChangeExt(src_path, ext)
	static_path = gbld_fs.AddPrefix(static_path, prefix)
	return gbld.NewFile(static_path)
}
