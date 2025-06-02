package gbld_fs

import (
	"github.com/Ankizle/gbld"
)

func TouchAll(f gbld.File) error {
	SetupFile(f)
	return Touch(f)
}
