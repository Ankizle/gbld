package gbld

import (
	"errors"
	"fmt"
	"path/filepath"
)

func (pj *Project) AddExecutable(name string, out string) *Module {
	mod := pj.AddModuleDefault(COMPILE_EXECUTABLE, name)
	mod.out_file = out
	return mod
}

func (mod *Module) CompileExecutable() error {
	// we've already compiled to .o files and dependencies
	// now just turn that into an executable

	o_paths := mod.get_object_paths()

	var ext string
	var rpath string
	switch mod.pj.OS {
	case "linux":
		ext = ""
		rpath = "-Wl,-rpath,$ORIGIN"
	case "windows":
		ext = ".exe"
		rpath = "-Wl,-rpath,."
	default:
		return errors.New("unsupported operating system: " + mod.pj.OS)
	}

	var args []string
	args = append(args, mod.pj.CC)
	args = append(args, mod.pj.FLAGS...)                                      // user flags
	args = append(args, "-o", filepath.Join(mod.pj.public, mod.out_file+ext)) // output file
	args = append(args, o_paths...)                                           // files to compile
	args = append(args, rpath)                                                // add rpath
	args = append(args, "-L"+mod.pj.public)                                   // directory to look for link targets
	args = append(args, mod.pj.get_libs()...)                                 // get libraries to link

	fmt.Println(args)

	cmd := mod.pj.Command(
		args,
		nil,
	)

	return to_compile_error(cmd.CombinedOutput())
}
