package optional

import (
	"testing"
)

func TestMin(t *testing.T) {
	s := &sliceIntOptional{v: []int{6, 3, 4, 5, 2}}
	min := s.Min()
	t.Log(min)
}
func TestMax(t *testing.T) {
	s := &sliceIntOptional{v: []int{6, 3, 4, 5, 8}}
	max := s.Max()
	t.Log(max)
}

