package gbld_c

import (
	"sync"

	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func DefaultBuildShared(pj *gbld.Project, mod *gbld.Module, filenames []string, out_name string) (out gbld.File) {
	// object files
	var objs []gbld.File
	var updated_objs []gbld.File

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

		// check if any dependencies were changed
		if gbld_fs.MaxTimestamp(deps) < gbld_fs.Timestamp(obj) {
			pj.Log("no work:", obj.Path())
			continue // no need to recompile
		} else {
			updated_objs = append(updated_objs, obj)
			pj.Log("building:", obj.Path())
			cmd.ExecAsync(&wg, mod.Abs("."), func(output []byte) {
				pj.LogErr(string(output))
			})
		}
	}

	wg.Wait() // wait for the objects to build

	// shared library
	out = Shared(pj.Getenv("OS"), pj.PublicAbs(out_name))

	pj.Log("building:", mod.Root())

	gbld_fs.SetupFile(out)

	cmd := mod.NewCommand()
	cmd.AddFlag("-shared", true)
	cmd.AddFlag("-o", out.Path())

	cmd.AddArgs(gbld_fs.Paths(objs)...)
	cmd.SetName(pj.Getenv("CPP"))

	cmd.Exec(mod.Abs("."), func(output []byte) {
		pj.LogErr(string(output))
	})

	return out
}
