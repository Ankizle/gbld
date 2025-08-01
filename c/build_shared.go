package gbld_c

import (
	"strings"

	"github.com/Ankizle/gbld"
	gbld_fs "github.com/Ankizle/gbld/fs"
)

func DefaultBuildShared(pj *gbld.Project, mod *gbld.Module, filenames []string, out_name string) (out gbld.File) {
	// object files
	mod.AddCommandFlag("-fPIC", true)
	objs, updated_objs := DefaultBuildObjects(pj, mod, filenames)

	// shared library
	out = Shared(pj.Getenv("OS"), pj.PublicAbs(out_name))

	if len(updated_objs) == 0 {
		pj.Log("no work:", mod.Root())
		return out
	}

	pj.Log("building:", mod.Root())

	gbld_fs.SetupFile(out)

	cmd := mod.NewCommand()
	cmd.AddFlag("-shared", true)
	cmd.AddFlag("-o", out.Path())

	cmd.AddArgs(gbld_fs.Paths(objs)...)
	cmd.SetName(pj.Getenv("CPP"))

	cmd.Exec(mod.Abs("."), func(output []byte) {
		pj.LogErr(
			"error while building", out.Path(),
			"\n"+strings.Join(cmd.GetArgList(), " "),
			"\n"+string(output),
		)
	})

	return out
}
