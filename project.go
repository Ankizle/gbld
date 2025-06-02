package gbld

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type Project struct {
	root   string
	build  string
	public string

	env           map[string]string
	command_flags map[string]interface{}

	log_file *os.File
	log_mtx  sync.Mutex

	modules []*Module
}

func NewProject(root string, build string, public string) *Project {
	pj := new(Project)
	pj.root = root
	pj.build = build
	pj.public = public
	pj.env = make(map[string]string)

	pj.log_file = os.Stdout // default

	// set a few environment variables
	pj.env["OS"] = runtime.GOOS
	pj.env["ARCH"] = runtime.GOARCH

	return pj
}

func NewProjectDefault() (*Project, error) {
	wd, e := os.Getwd()
	if e != nil {
		return nil, e
	}
	return NewProject(wd, filepath.Join(wd, "build"), filepath.Join(wd, "build/public")), nil
}

func (pj *Project) SetLogFile(f *os.File) {
	pj.log_file = f
}

func (pj *Project) Setenv(name string, value string) {
	pj.env[name] = value
}

func (pj *Project) Getenv(name string) string {
	return pj.env[name]
}

func (pj *Project) Log(v ...interface{}) {
	pj.log_mtx.Lock()
	fmt.Fprintln(pj.log_file, v...)
	pj.log_mtx.Unlock()
}
func (pj *Project) LogErr(v ...interface{}) {
	pj.log_mtx.Lock()
	fmt.Fprintln(pj.log_file, v...)
	pj.log_mtx.Unlock()
	os.Exit(1)
}

func (pj *Project) Abs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(pj.root, path)
}

func (pj *Project) BuildAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(pj.build, path)
}

func (pj *Project) PublicAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(pj.public, path)
}

func (pj *Project) Compile(wg *sync.WaitGroup) {
	for _, m := range pj.modules {
		m.CompileAsync(wg)
	}
	wg.Wait()
}
