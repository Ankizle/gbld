package gbld_fs

import "github.com/Ankizle/gbld"

func Paths(fs []gbld.File) []string {
	paths := make([]string, len(fs))
	for k, f := range fs {
		paths[k] = f.Path()
	}
	return paths
}
