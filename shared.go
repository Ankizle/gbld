package gbld

import (
	"errors"
	"path/filepath"
)

func (pj *Project) AddShared(name string, out string) *Module {
	mod := pj.AddModuleDefault(COMPILE_SHARED, name)
	mod.out_file = out
	return mod
}

func (mod *Module) CompileShared() error {
	// we've already compiled to .o files
	// just make a .so now

	o_paths := mod.get_object_paths()

	var ext string
	switch mod.pj.OS {
	case "linux":
		ext = ".so"
	case "windows":
		ext = ".dll"
	default:
		return errors.New("unsupported operating system: " + mod.pj.OS)
	}

	var args []string
	args = append(args, mod.pj.CC)
	args = append(args, mod.pj.FLAGS...)                                            // user flags
	args = append(args, "-shared", "-fPIC")                                         // make a shared library
	args = append(args, "-o", filepath.Join(mod.pj.public, "lib"+mod.out_file+ext)) // output file
	args = append(args, o_paths...)                                                 // files to compile

	cmd := mod.pj.Command(
		args,
		nil,
	)

	return to_compile_error(cmd.CombinedOutput())
}
