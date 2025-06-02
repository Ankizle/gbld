package gbld_fs

import (
	"os"
)

func SetupDir(path string) {
	os.MkdirAll(path, os.ModePerm)
}
