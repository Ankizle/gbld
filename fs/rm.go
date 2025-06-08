package gbld_fs

import "os"

// rm -rf
func RmPath(path string) {
	os.RemoveAll(path)
}
