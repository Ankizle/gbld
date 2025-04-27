package gbld

import (
	"crypto/md5"
	"os"
	"path/filepath"
)

type File struct {
	mod *Module

	name string
	path string
	out  string

	hash [16]byte
}

func (mod *Module) AddFile(name string) *File {
	file := &File{
		mod: mod,

		name: name,
		path: filepath.Join(mod.src, name),
		out:  filepath.Join(mod.out, change_ext(name, ".o")),
	}

	// read the file and hash it
	file_dat, e := os.ReadFile(file.path)
	if e == nil {
		file.hash = md5.Sum(file_dat)
	}

	mod.files = append(mod.files, file)
	return file
}

func (file *File) Compile() error {

	var args []string
	args = append(args, file.mod.pj.CC)
	args = append(args, file.mod.pj.FLAGS...)          // user defined flags
	args = append(args, file.mod.pj.get_includes()...) // get include directories
	args = append(args, "-o", file.out)                // output file
	args = append(args, "-c")                          // compile to .o file
	args = append(args, file.path)                     // specify the path of the file

	cmd := file.mod.pj.Command(
		args,
		file,
	)

	// check if we actually need to recompile
	if file.hash == file.mod.hashes[file.name] { // check if hash matches
		if _, e := os.Stat(file.out); e == nil { // check if .o file exists (user may have broken it)
			return nil // no need to recompile
		}
	}

	file.mod.hashes[file.name] = file.hash

	return to_compile_error(cmd.CombinedOutput())
}
