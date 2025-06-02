package gbld_fs

import (
	"math"

	"github.com/Ankizle/gbld"
)

func MaxTimestamp(fs []gbld.File) uint64 {

	if len(fs) == 0 {
		return math.MaxUint64
	}

	var max uint64 = 0
	for _, f := range fs {
		t := Timestamp(f)
		if t > max {
			max = t
		}
	}
	return max
}
