package optional

import "strconv"

type float64Optional struct {
	err error
	v   float64
}

func (o *float64Optional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *float64Optional) GetError() error {
	return o.err
}
func (o *float64Optional) Float64() float64 {
	return o.v
}
func (o *float64Optional) Float32() float32 {
	return float32(o.v)
}
func (o *float64Optional) String() string {
	return strconv.FormatFloat(o.v, 'E', -1, 64)
}
func (o *float64Optional) StringOptional() *stringOptional {
	return &stringOptional{err: o.err, v: strconv.FormatFloat(o.v, 'E', -1, 64)}
}
