package gbld

import (
	"os/exec"
	"path/filepath"
)

type File struct {
	mod *Module

	name string
	out  string
}

func (mod *Module) AddFile(name string) *File {
	file := &File{
		mod: mod,

		name: filepath.Join(mod.src, name),
		out:  filepath.Join(mod.out, change_ext(name, ".o")),
	}
	mod.files = append(mod.files, file)
	return file
}

func (file *File) Compile() error {

	var args []string
	args = append(args, file.mod.pj.FLAGS...) // user defined flags
	args = append(args, "-o", file.out)       // output file
	args = append(args, "-c")                 // compile to .o file

	exec.Command(
		file.mod.pj.CC,
		args...,
	)

	return nil
}
