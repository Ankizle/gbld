package gbld_spv

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func SourceHLSL(f string) gbld.File {
	// .hlsl shader
	return gbld.NewFile(gbld_fs.ChangeExt(f, ".hlsl"))
}
