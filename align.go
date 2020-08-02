package optional

type align struct {
	name string
	a    interface{}
}

func Align(name string, a interface{}) align {
	return align{
		name: name, a: a,
	}
}
