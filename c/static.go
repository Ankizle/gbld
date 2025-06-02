package gbld_c

import "github.com/Ankizle/gbld"

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

	static_path := change_extension(src_path, ext)
	static_path = add_prefix(static_path, prefix)
	return gbld.NewFile(static_path)
}
