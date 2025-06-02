package gbld_c

import "github.com/Ankizle/gbld"

func Executable(os string, f string) gbld.File {
	src_path := f
	var ext string

	switch os {
	case "linux":
		ext = ""
	case "windows":
		ext = ".exe"
	}

	shared_path := change_extension(src_path, ext)
	return gbld.NewFile(shared_path)
}
