package gbld_c

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func DefaultBuildObjects(pj *gbld.Project, mod *gbld.Module, filenames []string) (objs []gbld.File, updated_objs []gbld.File) {
	// object files

	var wg sync.WaitGroup

	for _, f := range filenames {

		// files we'll need
		obj := Object(mod.BuildAbs(f))
		src := SourceCPP(mod.Abs(f))
		deps := Deps(mod.BuildAbs(f))

		objs = append(objs, obj)

		// create the object file if it doesn't exist
		gbld_fs.SetupFile(obj)

		// execute the build command
		cmd := mod.NewCommand()
		cmd.AddFlag("-MMD", true)

		cmd.AddFlag("-c", true)
		cmd.AddFlag("-fPIC", true)
		cmd.AddFlag("-o", obj.Path())

		cmd.AddArg(src.Path())
		cmd.SetFile(src)

		cmd.SetName(pj.Getenv("CPP"))

		command_string := []byte(fmt.Sprint(cmd.GetArgList()))
		command_string_file, e_open := os.OpenFile(SourceCommandStringFile(obj.Path()).Path(), os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer command_string_file.Close()

		command_string_file_string := make([]byte, len(command_string))
		n_read, e_read := command_string_file.Read(command_string_file_string)

		// for the old command string and new command string to be equal, they need to have the same length
		command_strings_are_equal := n_read == len(command_string)
		if command_strings_are_equal {
			// if they have the same length, then carry on checking if all their bytes match
			for i := range command_string {
				if command_string[i] != command_string_file_string[i] {
					command_strings_are_equal = false
					break
				}
			}
		}

		// check if any dependencies were changed
		if gbld_fs.MaxTimestamp(deps) < gbld_fs.Timestamp(obj) && /* previous build of this file is newer than update */
			e_open == nil && e_read == nil && /* command string file is not missing */
			command_strings_are_equal /* the commands match */ {
			pj.Log("no work:", obj.Path())
			continue // no need to recompile
		} else {
			updated_objs = append(updated_objs, obj)
			pj.Log("building:", obj.Path())
			cmd.ExecAsync(&wg, mod.Abs("."), func(output []byte) {
				pj.LogErr(
					"error while building", obj.Path(),
					"\n"+strings.Join(cmd.GetArgList(), " "),
					"\n"+string(output),
				)
			})

			// log the command used to build this file
			command_string_file.Seek(0, 0)
			command_string_file.Write(command_string)
			command_string_file.Truncate(int64(len(command_string)))
		}
	}

	wg.Wait() // wait for the objects to build

	return objs, updated_objs
}
