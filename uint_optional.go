package optional


type uintOptional struct {
	err error
	v   uint
}

func (o *uintOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *uintOptional) GetError() error {
	return o.err
}
func (o *uintOptional) Uint() uint {
	return o.v
}
