package optional

import (
	"errors"
	"fmt"
)

type optional struct {
	err error
	v   interface{}
}

func (o *optional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *optional) GetError() error {
	return o.err
}

func (o *optional) StringOptional() *stringOptional {
	if o.err != nil {
		return &stringOptional{
			err: o.err,
		}
	}
	s, ok := o.v.(string)
	if ok {
		return &stringOptional{
			v: s,
		}
	}
	return &stringOptional{
		v:   s,
		err: errors.New("类型转换失败"),
	}
}

func (o *optional) String() string {
	return fmt.Sprint(o.v)
}
func Optional(v interface{}) *optional {
	return &optional{v: v}
}

func BoolOptional(b bool) *boolOptional {
	return &boolOptional{v: b}
}
func StringOptional(v string) *stringOptional {
	return &stringOptional{v: v}
}
func IntOptional(v int) *intOptional {
	return &intOptional{v: v}
}
func UintOptional(v uint) *uintOptional {
	return &uintOptional{v: v}
}
func Int64Optional(v int64) *int64Optional {
	return &int64Optional{v: v}
}
func SliceIntOptional(v ...int) *sliceIntOptional {
	return &sliceIntOptional{v: v}
}
func SliceIntOptionals(v []int) *sliceIntOptional {
	return &sliceIntOptional{v: v}
}
func SliceStringOptional(v ...string) *sliceStringOptional {
	return &sliceStringOptional{v: v}
}
func SliceStringOptionals(v []string) *sliceStringOptional {
	return &sliceStringOptional{v: v}
}
