package gbld

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Project struct {
	CC         string
	FLAGS      []string
	OS         string
	PERMISSION os.FileMode

	name string

	root   string
	out    string
	public string

	includes []string

	cmd_log []CommandLog

	externals []*External // external deps
	modules   []*Module   // local deps
}

func NewProject(name string, root string, out string, public string) *Project {
	pj := &Project{}
	pj.name = name
	pj.root = root
	pj.out = filepath.Join(pj.root, out)
	pj.public = filepath.Join(pj.out, public)
	pj.PERMISSION = 05777
	return pj
}

func NewProjectDefault(name string) *Project {
	pj := NewProject(name, ".", "build", "public")
	pj.CC = "clang++"
	pj.FLAGS = []string{"-fdiagnostics-color=always"}
	pj.OS = runtime.GOOS
	return pj
}

func (pj *Project) get_libs() (libs []string) {

	for _, ext := range pj.externals {
		libs = append(libs, ext.get_libs()...) // link external libs
	}

	for _, mod := range pj.modules {
		if mod.compile_mode == COMPILE_SHARED || mod.compile_mode == COMPILE_STATIC {
			libs = append(libs, "-l"+mod.name) // link libraries
		}
	}
	return libs
}

func (pj *Project) get_includes() []string {
	return pj.includes
}

func (pj *Project) Include(path string) {
	pj.includes = append(pj.includes, "-I"+path)
}

func (pj *Project) Command(args []string, file *File) *exec.Cmd {
	pj.cmd_log = append(pj.cmd_log, CommandLog{
		args: args,
		file: file,
	})
	return exec.Command(args[0], args[1:]...)
}

func (pj *Project) Compile() error {

	// make sure the target and public directories exit
	os.MkdirAll(pj.out, pj.PERMISSION)
	os.MkdirAll(pj.public, pj.PERMISSION)

	// compile each external dependency
	for _, ext := range pj.externals {
		if e := ext.Compile(); e != nil {
			return e // encountered an error in the dependency
		}
	}

	// compile each module
	for _, mod := range pj.modules {
		if e := mod.Compile(); e != nil {
			return e // encountered a compile-time error
		}
	}
	return nil
}
