package gbld_c

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

// gbld_c.Object(mod.BuildAbs(f))
// 			deps := gbld_c.Deps(mod.BuildAbs(f))
// 			src := gbld_c.SourceCPP(mod.Abs(f))

// take a source file name and return its object file name (file.cpp -> file.o)
func Object(f string) gbld.File {
	src_path := f
	obj_path := gbld_fs.ChangeExt(src_path, ".o")
	return gbld.NewFile(obj_path)
}
