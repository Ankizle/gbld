package gbld_c

import "github.com/Ankizle/gbld"

func Shared(os string, f string) gbld.File {
	src_path := f
	var ext string

	switch os {
	case "linux":
		ext = ".so"
	case "windows":
		ext = ".dll"
	}

	shared_path := change_extension(src_path, ext)
	return gbld.NewFile(shared_path)
}
