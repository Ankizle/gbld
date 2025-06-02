package gbld_fs

import (
	"os"

	"github.com/Ankizle/gbld"
)

func Timestamp(f gbld.File) uint64 {
	stat, e := os.Stat(f.Path())

	if e != nil {
		return 0
	}

	return uint64(stat.ModTime().Unix())
}
