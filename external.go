package gbld

type External struct {
	origin string
}

func NewExternal(origin string) *External {
	return &External{
		origin: origin,
	}
}
