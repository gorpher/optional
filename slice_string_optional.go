package optional

import "strings"

type sliceStringOptional struct {
	err error
	v   []string
}

func (o *sliceStringOptional) HashError() bool {
	if o.err != nil {
		return true
	}
	return false
}
func (o *sliceStringOptional) GetError() error {
	return o.err
}
func (o *sliceStringOptional) Slice() []string {
	return o.v
}

func (o *sliceStringOptional) Map(mp func(s string) string) *sliceStringOptional {
	for i := range o.v {
		o.v[i] = mp(o.v[i])
	}
	return o
}

func (o *sliceStringOptional) ToLower() *sliceStringOptional {
	return o.Map(strings.ToLower)
}

func (o *sliceStringOptional) ToUpper() *sliceStringOptional {
	return o.Map(strings.ToUpper)
}