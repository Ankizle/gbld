package gbld_c

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func SourceC(f string) gbld.File {
	src_path := f
	new_src_path := gbld_fs.ChangeExt(src_path, ".c")
	return gbld.NewFile(new_src_path)
}

func SourceCPP(f string) gbld.File {
	src_path := f
	new_src_path := gbld_fs.ChangeExt(src_path, ".cpp")
	return gbld.NewFile(new_src_path)
}

func SourceCC(f string) gbld.File {
	src_path := f
	new_src_path := gbld_fs.ChangeExt(src_path, ".cc")
	return gbld.NewFile(new_src_path)
}

func SourceCXX(f string) gbld.File {
	src_path := f
	new_src_path := gbld_fs.ChangeExt(src_path, ".cxx")
	return gbld.NewFile(new_src_path)
}

func SourceCommandStringFile(f string) gbld.File {
	// this is a special file that notes the command used to build a .o file
	// extension is .csf
	src_path := f
	csf := gbld_fs.ChangeExt(src_path, ".csf")
	return gbld.NewFile(csf)
}
