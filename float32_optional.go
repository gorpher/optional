package optional

import "strconv"

type float32Optional struct {
	err error
	v   float32
}


func (o *float32Optional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *float32Optional) GetError() error {
	return o.err
}
func (o *float32Optional) Float32() float32 {
	return o.v
}
func (o *float32Optional) String() string {
	return strconv.FormatFloat(float64(o.v), 'E', -1, 32)
}
func (o *float32Optional) StringOption() *stringOptional {
	return &stringOptional{err: o.err, v: strconv.FormatFloat(float64(o.v), 'E', -1, 32)}
}
