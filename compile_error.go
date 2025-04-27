package gbld

type CompileError struct {
	msg []byte
}

func to_compile_error(msg []byte, e error) error {
	if e != nil {
		return &CompileError{msg}
	} else {
		return nil
	}
}

func (e *CompileError) Error() string {
	return string(e.msg)
}
