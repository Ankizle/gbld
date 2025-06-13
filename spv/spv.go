package gbld_spv

import (
	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func SourceSPV(f string) gbld.File {
	// change .glsl, .hlsl, or any other shader extension to just .spv
	return gbld.NewFile(gbld_fs.ChangeExt(f, ".spv"))
}
