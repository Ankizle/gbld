package gbld_fs

import (
	"path/filepath"

	"github.com/Ankizle/gbld"
)

func SetupFile(f gbld.File) {
	dir := filepath.Dir(f.Path())
	SetupDir(dir)
}
