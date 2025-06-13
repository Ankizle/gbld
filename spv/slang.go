package gbld_spv

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func SourceSlang(f string) gbld.File {
	// .slang shader
	return gbld.NewFile(gbld_fs.ChangeExt(f, ".slang"))
}
