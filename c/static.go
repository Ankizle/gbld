package gbld_c

import "github.com/Ankizle/gbld"

func Static(os string, f string) gbld.File {
	src_path := f
	var ext string

	switch os {
	case "linux":
		ext = ".a"
	case "windows":
		ext = ".lib"
	}

	shared_path := change_extension(src_path, ext)
	return gbld.NewFile(shared_path)
}
