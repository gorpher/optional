package optional

import "strconv"

type int64Optional struct {
	err error
	v   int64
}


func (o *int64Optional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *int64Optional) GetError() error {
	return o.err
}
func (o *int64Optional) Int64() int64 {
	return o.v
}
func (o *int64Optional) String() string{
	return  strconv.FormatInt(o.v, 10)
}
func (o *int64Optional) StringOptional() *stringOptional {
	return &stringOptional{err: o.err, v: strconv.FormatInt(o.v, 10)}
}
