package gbld_fs

import (
	"os"
	"path/filepath"

	"github.com/Ankizle/gbld"
)

func SetupFile(f gbld.File) {
	dir := filepath.Dir(f.Path())
	os.MkdirAll(dir, os.ModePerm)
}
