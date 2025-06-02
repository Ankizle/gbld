package gbld_c

import (
	"os"
	"strings"

	"github.com/Ankizle/gbld"
)

// inputs a source file path
// outputs a list of dependencies obtained from the source file's .d file
func Deps(f string) []gbld.File {

	// the general structure of .d files we'll see are

	// object.o: dependency1.[cpp/h] dependency2.[cpp/h] ...
	// there may be backslashes between dependencies

	// first convert f (which is given as .cpp or .c) to .d
	dep_path := change_extension(f, ".d")
	dep_src_bytes, e := os.ReadFile(dep_path)
	dep_src := string(dep_src_bytes)

	if e != nil {
		return []gbld.File{}
	}

	// otherwise, read all the dependencies
	// this is a bit hacky (may not be fully compliant with the standard), TODO

	dep_list_src := strings.TrimSpace(strings.Split(dep_src, ":")[1])
	dep_list_raw := strings.Split(dep_list_src, " ")
	var dep_list []gbld.File

	for _, dep := range dep_list_raw {
		if dep == "" || dep == "\\" {
			continue // ignore backslashes and other things like that
		}

		dep_list = append(dep_list, gbld.NewFile(dep)) // only add the source/header files
	}

	return dep_list
}
