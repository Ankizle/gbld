package gbld_fs

import (
	"github.com/Ankizle/gbld"
)

func MaxTimestamp(fs []gbld.File) uint64 {
	var max uint64 = 0
	for _, f := range fs {
		t := Timestamp(f)
		if t > max {
			max = t
		}
	}
	return max
}
