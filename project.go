package gbld

import (
	"path/filepath"
	"runtime"
)

type Project struct {
	CC    string
	FLAGS []string
	OS    string

	name string

	root   string
	out    string
	public string

	modules []*Module
}

func NewProject(name string, root string, out string, public string) *Project {
	pj := &Project{}
	pj.name = name
	pj.root = root
	pj.out = filepath.Join(pj.root, out)
	pj.public = filepath.Join(pj.root, public)
	return pj
}

func NewProjectDefault(name string) *Project {
	pj := NewProject(name, ".", "target", "public")
	pj.CC = "clang++"
	pj.FLAGS = []string{"-fdiagnostics-color=always"}
	pj.OS = runtime.GOOS
	return pj
}

func (pj *Project) Compile() error {
	for _, mod := range pj.modules {
		if e := mod.Compile(); e != nil {
			return e // encountered a compile-time error
		}
	}
	return nil
}
