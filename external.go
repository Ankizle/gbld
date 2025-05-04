package gbld

import (
	"os"
	"path/filepath"
	"strings"
)

type ExternalData struct {
	BuildCommands [][]string

	Name    string
	License string
	Bin     string
	Libs    []string
}

type External struct {
	pj   *Project
	data ExternalData
}

func (pj *Project) AddExternal(data ExternalData) *External {
	ext := &External{}
	ext.pj = pj
	ext.data = data

	pj.externals = append(pj.externals, ext)
	return ext
}

func (ext *External) get_libs() (libs []string) {
	if ext.data.Bin != "" {
		for _, lib := range ext.data.Libs {
			libs = append(libs, "-l"+lib) // add the libs to use
		}
		libs = append(libs, "-L"+ext.data.Bin) // linking directory (compile time)
	}
	return libs
}

func (ext *External) Compile() error {

	// run the build commands
	// TODO: make this code cleaner/rework to be more safe
	var cmds []string
	for _, cmd_code := range ext.data.BuildCommands {
		cmds = append(cmds, strings.Join(cmd_code, " "))
	}

	cmds_str := strings.Join(cmds, "&&")

	if len(cmds) > 0 {
		cmd := ext.pj.Command([]string{"/bin/sh", "-c", cmds_str}, nil)
		out, e := cmd.CombinedOutput()
		if e := to_compile_error(out, e); e != nil {
			return e
		}
	}

	// copy the binaries and license to the public output directory

	// binaries
	for _, lib := range ext.data.Libs {
		lib_data, filename, e := read_all_ext(ext.data.Bin, "lib"+lib, []string{".so.3", ".so.1", ".so", ".a"})

		if e != nil {
			return e
		}

		os.WriteFile(
			filepath.Join(ext.pj.public, filename),
			lib_data,
			ext.pj.PERMISSION,
		)
	}

	// license
	license_dir := filepath.Join(ext.pj.public, "licenses")
	os.MkdirAll(license_dir, ext.pj.PERMISSION)
	license_dat, e := os.ReadFile(ext.data.License)
	if e != nil {
		return e
	}
	e = os.WriteFile(
		filepath.Join(license_dir, ext.data.Name), // copy to build/public/LICENSE
		license_dat,
		ext.pj.PERMISSION,
	)
	if e != nil {
		return e
	}

	return nil
}
