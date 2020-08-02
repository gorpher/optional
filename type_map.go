package optional

import (
	"bytes"
	"errors"
	"fmt"
)

type typeStringMap struct {
	typeImplSigil
	AttrType map[string]Type
}

func StringMap() Type {
	return Type{
		typeStringMap{},
	}
}

func StringMapType(ty map[string]Type) Type {
	return Type{
		typeStringMap{
			AttrType: ty,
		},
	}
}

func (t typeStringMap) Equals(other Type) bool {
	_, isMap := other.typeImpl.(typeStringMap)
	return isMap
}
func (t typeStringMap) FriendlyName() string {
	b := bytes.NewBufferString("map of ")
	b.WriteString("[")
	for k, v := range t.AttrType {
		b.WriteString("(")
		b.WriteString(k + "=" + v.FriendlyName())
		b.WriteString(")")
	}
	b.WriteString("]")
	return b.String()
}
func (t typeStringMap) GoString() string {
	return fmt.Sprintf("optional.Map(%#value)", t.AttrType)
}

func (t typeStringMap) UpdateAttrType(ty map[string]Type) {
	for k := range ty {
		t.AttrType[k] = ty[k]
	}
}
func (t typeStringMap) GetStringMapType(key string) Type {
	return t.AttrType[key]
}

func (t Type) IsMapType() bool {
	_, ok := t.typeImpl.(typeStringMap)
	return ok
}
func (t Type) UpdateAttrType(ty map[string]Type) error {
	tm, ok := t.typeImpl.(typeStringMap)
	if ok {
		tm.UpdateAttrType(ty)
	}
	return errors.New("type not support ")
}
func (t Type) GetStringMapType(key string) (Type, error) {
	tm, ok := t.typeImpl.(typeStringMap)
	if ok {
		return tm.GetStringMapType(key), nil
	}
	return Type{}, errors.New("type not support ")
}

func (t Type) Len() int {
	if lt, ok := t.typeImpl.(typeStringMap); ok {
		return len(lt.AttrType)
	}
	return 0
}
