package optional

import "strconv"

type intOptional struct {
	err error
	v   int
}


func (o *intOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *intOptional) GetError() error {
	return o.err
}
func (o *intOptional) Int() int {
	return o.v
}
func (o *intOptional) Int64() int64 {
	return int64(o.v)
}
func (o *intOptional) String() string {
	return strconv.Itoa(o.v)
}
func (o *intOptional) Int64Optional() *int64Optional {
	return &int64Optional{v: int64(o.v), err: o.err}
}
func (o *intOptional) StringOptional() *stringOptional {
	// strconv.FormatInt(int64(int), 10)
	return &stringOptional{v: strconv.Itoa(o.v), err: o.err}
}
