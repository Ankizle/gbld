package gbld

import "path/filepath"

type Module struct {
	pj *Project

	name string

	src string
	out string

	files []*File
}

func (pj *Project) AddModule(name string, src string, out string) *Module {
	mod := &Module{
		pj: pj,

		name: name,
		src:  src,
		out:  out,
	}
	pj.modules = append(pj.modules, mod)
	return mod
}

func (pj *Project) AddModuleDefault(name string) *Module {
	return pj.AddModule(name, filepath.Join(pj.root, name, "src"), filepath.Join(pj.out, name))
}

func (mod *Module) Compile() error {
	for _, file := range mod.files {
		if e := file.Compile(); e != nil {
			return e // error compiling a file to an object
		}
	}
	return nil
}
