package gbld

type CompileError struct {
	msg []byte
}

func NewCompileError(msg []byte, e error) *CompileError {
	if e != nil {
		return &CompileError{msg}
	} else {
		return nil
	}
}

func (e *CompileError) Error() string {
	return string(e.msg)
}
