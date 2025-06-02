package gbld_fs

import (
	"os"
	"time"

	"github.com/Ankizle/gbld"
)

func Touch(f gbld.File) (e error) {
	now := time.Now()
	e = os.Chtimes(f.Path(), now, now)
	if os.IsNotExist(e) {
		f, err := os.Create(f.Path())
		if err != nil {
			return err
		}
		return f.Close()
	}
	return e
}
