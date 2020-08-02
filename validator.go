package optional

import (
	"fmt"
)

type validator struct {
	matches []Match
	name    string // 字段名
	value   Value
	strict  bool // 严格模式
}

func (o validator) is(fn func(r rune) bool) bool {
	for _, v := range o.value.v.(string) {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (o validator) has(fn func(r rune) bool) bool {
	for _, v := range o.value.v.(string) {
		if fn(v) {
			return true
		}
	}
	return false
}
func (o validator) GetError() error {
	return o.Value().GetError()
}
func (o validator) Value() Value {
	if err := o.value.GetError(); err != nil {
		return o.value
	}
	for i := range o.matches {
		if err := o.matches[i](&o); err != nil {
			o.value.err = err
			return o.value
		}
	}
	return o.value
}
func (o validator) Align(a interface{}) error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	for i := range o.matches {
		if err := o.matches[i](&o); err != nil {
			return err
		}
	}
	return o.value.UnMarshal(a)
}
func (o validator) Processor(name string, apply ...Apply) processor {
	return o.value.Processor(name, apply...)
}

type validators struct {
	value  Value
	values map[string]validator
}

func (o validators) GetError() error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	for i := range o.values {
		if err := o.values[i].GetError(); err != nil {
			return err
		}
	}
	return nil
}
func (o validators) Value() Value {
	if err := o.value.GetError(); err != nil {
		return o.value
	}
	for _, v := range o.values {
		if err := v.Value().GetError(); err != nil {
			o.value.err = err
			return o.value
		}
	}
	return o.value
}
func (o validators) Align(a align) error {
	v := o.Value()
	if err := v.GetError(); err != nil {
		return err
	}
	v2, ok := o.values[a.name]
	if ok {
		return v2.Value().UnMarshal(a.a)
	}
	return fmt.Errorf("align variable %s value error", a.name)
}
func (o validators) Aligns(aligns ...align) error {
	if err := o.GetError(); err != nil {
		return err
	}
	for i := range aligns {
		key := aligns[i].name
		v2, ok := o.values[key]
		if ok {
			if err := v2.GetError(); err != nil {
				return err
			}
			o.value.GetMapValue(key).UnMarshal(aligns[i].a)
		}
		return fmt.Errorf("align variable %s value error", key)
	}
	return nil
}
func (o validators) Processors(ps ...processor) processors {
	return o.value.Processors(ps...)
}

type Match func(val *validator) error

func Validate(name string, matches ...Match) validator {
	return validator{
		name:    name,
		matches: matches,
	}
}
