package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Ankizle/gbld"
	gbld_c "github.com/Ankizle/gbld/c"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func main() {

	pj, e := gbld.NewProjectDefault()
	pj.SetLogFile(os.Stdout)

	pj.Setenv("CPP", "clang++")
	pj.Setenv("AR", "ar")

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	var files = []string{
		"impl/foo.cpp",
		"impl/bar.cpp",
		"impl/main.cpp",
	}

	mod := pj.AddModule("main")
	mod.SetCompileCallback(func() {

		// object files
		var objs []gbld.File

		var wg sync.WaitGroup

		for _, f := range files {

			// files we'll need
			obj := gbld_c.Object(mod.BuildAbs(f))
			src := gbld_c.SourceCPP(mod.Abs(f))
			deps := gbld_c.Deps(mod.BuildAbs(f), src)

			objs = append(objs, obj)

			// check if any dependencies were changed
			if gbld_fs.MaxTimestamp(deps) < gbld_fs.Timestamp(obj) {
				pj.Log("no work", obj.Path())
				continue // no need to recompile
			}

			pj.Log("building", obj.Path())

			// create the object file if it doesn't exist
			gbld_fs.SetupFile(obj)

			// execute the build command
			cmd := mod.NewCommand()
			cmd.SetFlag("-I", mod.Abs("."))
			cmd.SetFlag("-c", true)
			cmd.SetFlag("-fPIC", true)
			cmd.SetFlag("-o", obj.Path())

			cmd.SetArg(src.Path())
			cmd.SetFile(src)

			cmd.SetName(pj.Getenv("CPP"))

			cmd.ExecAsync(&wg, mod.Abs("."), func(output []byte) {
				pj.LogErr(string(output))
			})
		}

		// wait for all the objects to build
		wg.Wait()

		// shared library
		out := gbld_c.Executable(pj.Getenv("OS"), pj.PublicAbs("example_00"))

		if gbld_fs.MaxTimestamp(objs) < gbld_fs.Timestamp(out) {
			pj.Log("no work", mod.Root())
			return // no need to recompile
		}

		pj.Log("building", mod.Root())

		// create the executable file if it doesn't exist
		gbld_fs.SetupFile(out)

		cmd := mod.NewCommand()
		cmd.SetFlag("-o", out.Path())

		cmd.SetArgs(gbld_fs.Paths(objs)...)
		cmd.SetName(pj.Getenv("CPP"))

		fmt.Println(cmd.GetArgList())

		cmd.Exec(mod.Abs("."), func(output []byte) {
			pj.LogErr(string(output))
		})
	})

	var wg sync.WaitGroup
	pj.Compile(&wg)
}
