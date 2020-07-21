package optional

type sliceIntOptional struct {
	err error
	v   []int
}

func (o *sliceIntOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *sliceIntOptional) GetError() error {
	return o.err
}
func (o *sliceIntOptional) Slice() []int {
	return o.v
}

func (o *sliceIntOptional) Min() int {
	var i = o.v[0]
	for _, v := range o.v {
		if v < i {
			i = v
		}
	}
	return i
}
func (o *sliceIntOptional) Max() int {
	var i = o.v[0]
	for _, v := range o.v {
		if v > i {
			i = v
		}
	}
	return i
}
