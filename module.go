package gbld

import (
	"fmt"
	"path/filepath"
	"sync"
)

type Module struct {
	pj *Project

	root  string
	build string

	command_flags map[string]interface{}

	compile_callback func()

	command_log []*Command
}

func (pj *Project) AddModule(relative_path string) *Module {

	root := filepath.Join(pj.root, relative_path)
	build := filepath.Join(pj.build, relative_path)

	mod := new(Module)
	mod.pj = pj
	mod.root = root
	mod.build = build

	mod.command_flags = make(map[string]interface{})

	mod.compile_callback = func() {
		fmt.Fprintln(pj.log_file, "warning:", "compile callback for", relative_path, "is unset")
	}

	pj.modules = append(pj.modules, mod)

	return mod
}

func (mod *Module) Root() string {
	return mod.root
}

func (mod *Module) SetCompileCallback(cb func()) {
	mod.compile_callback = cb
}

func (mod *Module) Abs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(mod.root, path)
}

func (mod *Module) BuildAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(mod.build, path)
}

func (mod *Module) CompileAsync(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mod.compile_callback()
		wg.Done()
	}()
}
