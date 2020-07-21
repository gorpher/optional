package optional


type uint64Optional struct {
	err error
	v   uint64
}


func (o *uint64Optional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *uint64Optional) GetError() error {
	return o.err
}
func (o *uint64Optional) uInt64() uint64 {
	return o.v
}
