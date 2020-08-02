package optional

import (
	"fmt"
)

type processor struct {
	name    string
	applies []Apply
	value   Value
}

func (o processor) GetError() error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	return nil
}

func (o processor) Value() Value {
	if err := o.value.GetError(); err != nil {
		return o.value
	}
	for i := range o.applies {
		if err := o.applies[i](&o); err != nil {
			o.value.err = err
			return o.value
		}
	}
	return o.value
}

func (o processor) Align(a align) error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	if err := o.value.UnMarshal(a.a); err != nil {
		return errorf("%s filed align error : %value", a.name, err)
	}
	return fmt.Errorf("align variable %s value error", a.name)
}

func (o processors) Validate(name string, match ...Match) validator {
	return o.value.Validate(name, match...)
}

type processors struct {
	value  Value
	values map[string]processor
}

func (o processors) GetError() error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	if err := o.Value().GetError(); err != nil {
		return err
	}
	return nil
}

func (o processors) Value() Value {
	if err := o.value.GetError(); err != nil {
		return o.value
	}
	for i := range o.values {
		value := o.values[i].Value()
		if err := value.GetError(); err != nil {
			return value
		}
		//todo
		o.value.SetMapValue(o.values[i].name, value)
	}
	return o.value
}

func (o processors) Aligns(aligns ...align) error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	return o.value.Aligns(aligns...)
}

func (o processors) Align(a align) error {
	if err := o.value.GetError(); err != nil {
		return err
	}
	return o.value.Aligns(a)
}

func (o processors) Validates(validates ...validator) validators {
	return o.value.Validates(validates...)
}

type Apply func(val *processor) error

func Process(name string, apply ...Apply) processor {
	return processor{
		name:    name,
		applies: apply,
	}
}
