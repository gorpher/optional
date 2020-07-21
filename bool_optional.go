package optional

type boolOptional struct {
	err error
	v   bool
}

func (o *boolOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *boolOptional) GetError() error {
	return o.err
}

func (o *boolOptional) Bool() bool {
	return o.v
}
func (o *boolOptional) String() string {
	if o.v {
		return "true"
	}
	return "false"
}
func (o *boolOptional) Int() int {
	if o.v {
		return 1
	}
	return 0
}

// bool 类型的指针
func (o *boolOptional) SetValue(v interface{}) error {
	if o.err != nil {
		return o.err
	}
	//todo 设置值
	return nil

}
func (o *boolOptional) SetBool(v *bool) error {
	if o.err != nil {
		return o.err
	}
	*v = o.v
	return nil
}
