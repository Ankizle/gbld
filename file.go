package gbld

type File struct {
	path string
}

func NewFile(path string) File {
	return File{path}
}

func (f File) Path() string {
	return f.path
}

func (f File) IsReal() bool {
	return f.Path() != ""
}
