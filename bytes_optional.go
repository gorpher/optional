package optional

type bytesOptional struct {
	err error
	v   []byte
}


func (o *bytesOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *bytesOptional) GetError() error {
	return o.err
}

