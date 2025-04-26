package gbld

type Project struct {
	CC       string
	FLAGS    []string
	PLATFORM string

	name string

	root string
	out  string

	modules []*Module
}

func NewProject(name string, root string, out string) *Project {
	return &Project{
		name: name,
		root: root,
		out:  out,
	}
}

func NewProjectDefault(name string) *Project {
	pj := NewProject(name, ".", "./target")
	pj.CC = "clang++"
	pj.FLAGS = []string{"-fdiagnostics-color=always"}
	return pj
}

func (pj *Project) Compile() error {
	for _, mod := range pj.modules {
		if e := mod.Compile(); e != nil {
			return e // encountered a compile-time error
		}
	}
	return nil
}
