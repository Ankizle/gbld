package gbld_c

import "github.com/Ankizle/gbld"

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

	shared_path := change_extension(src_path, ext)
	shared_path = add_prefix(shared_path, prefix)
	return gbld.NewFile(shared_path)
}
