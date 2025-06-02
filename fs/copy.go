package gbld_fs

import (
	"io"
	"os"

	"github.com/Ankizle/gbld"
)

func Copy(src gbld.File, dst gbld.File) error {
	src_file, e := os.Open(src.Path())
	if e != nil {
		return e
	}
	defer src_file.Close()

	SetupFile(dst)
	dst_file, e := os.Create(dst.Path())
	if e != nil {
		return e
	}
	defer dst_file.Close()

	_, e = io.Copy(dst_file, src_file)
	if e != nil {
		return e
	}

	return nil
}
