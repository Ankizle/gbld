package gbld

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Module struct {
	pj           *Project
	compile_mode CompileMode

	name string

	src string
	out string

	out_file string // used when compiling static, shared, or executable

	files       []*File
	hashes_path string
	hashes      map[string][16]byte
}

func (pj *Project) AddModule(mode CompileMode, name string, src string, out string) *Module {
	mod := &Module{
		pj:           pj,
		compile_mode: mode,

		name: name,
		src:  src,
		out:  out,
	}

	// read the file hashes
	mod.hashes_path = filepath.Join(mod.out, "hashes.json")
	hashes_data, _ := os.ReadFile(mod.hashes_path)
	e := json.Unmarshal(hashes_data, &mod.hashes)

	if e != nil {
		mod.hashes = make(map[string][16]byte)
	}

	pj.modules = append(pj.modules, mod)
	return mod
}

func (pj *Project) AddModuleDefault(mode CompileMode, name string) *Module {
	return pj.AddModule(mode, name, filepath.Join(pj.root, name, "src"), filepath.Join(pj.out, name))
}

func (mod *Module) get_object_paths() (o_paths []string) {
	for _, file := range mod.files {
		o_paths = append(o_paths, file.out)
	}
	return o_paths
}

func (mod *Module) Compile() error {

	// make sure the module's output directory exists
	os.MkdirAll(mod.out, mod.pj.PERMISSION)

	for _, file := range mod.files {
		if e := file.Compile(); e != nil {
			return e // error compiling a file to an object
		}
	}

	hashes_data, _ := json.Marshal(mod.hashes)
	if e := os.WriteFile(mod.hashes_path, hashes_data, mod.pj.PERMISSION); e != nil {
		return e
	}

	// now compile it to an executable, shared library, or static library

	switch mod.compile_mode {
	case COMPILE_EXECUTABLE:
		return mod.CompileExecutable()
	case COMPILE_SHARED:
		return mod.CompileShared()
	// case COMPILE_STATIC:
	default:
		return errors.New("unsupported compile mode")
	}
}
