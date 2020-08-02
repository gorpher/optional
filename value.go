package optional

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type Belief interface {
	GetError() error

	Value() Value
}

const (
	strict = true
)

type Value struct {
	err error
	ty  Type
	v   interface{}
}

func (val Value) Type() Type {
	return val.ty
}

func (val Value) IsNull() bool {
	return val.v == nil
}

func (val Value) Converter() converter {
	return converter{
		value: val,
	}
}

func (val Value) Equals(v Value) bool {
	if val.ty.Equals(v.ty) && val.v == val.v {
		return true
	}
	return false
}

func (val Value) GetError() error {
	return val.err
}

func (val Value) Value() Value {
	return val
}

func (val Value) IsPrimitiveValue() bool {
	return val.ty.IsPrimitiveType()
}

func (val Value) IsMapValue() bool {
	return val.ty.IsMapType()
}

var NilVal = Value{
	ty: Type{typeImpl: nil},
	v:  nil,
}

func (val Value) UnMarshal(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	if val.IsPrimitiveValue() {
		rve := rv.Elem()
		switch rve.Type().Kind() {
		case reflect.String:
			cv, err := val.Converter().String()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
			return nil
		case reflect.Int:
			cv, err := val.Converter().Int()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Int8:
			cv, err := val.Converter().Int8()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Int16:
			cv, err := val.Converter().Int16()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Int32:
			cv, err := val.Converter().Int32()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Int64:
			cv, err := val.Converter().Int64()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Uint:
			cv, err := val.Converter().Uint()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Uint8:
			cv, err := val.Converter().Uint8()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Uint16:
			cv, err := val.Converter().Uint16()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Uint32:
			cv, err := val.Converter().Uint32()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Uint64:
			cv, err := val.Converter().Uint64()
			if err != nil {
				return err
			}
			rve.Set(reflect.ValueOf(cv))
		case reflect.Ptr:
			v = &val.v
		}
		return nil
	}
	if val.IsMapValue() {
		//todo
	}
	return nil
}

func (val Value) StringProcessor() stringProcessor {
	return &linkProcessor{
		v: val,
	}
}

func (val Value) Validate(name string, match ...Match) validator {
	return validator{matches: match, name: name, value: val}
}

func (val Value) Validates(validates ...validator) validators {
	if val.IsMapValue() && val.ty.Len() == 0 {
		val.err = errorf("map value  is null!!!")
		return validators{
			value: val,
		}
	}
	m := make(map[string]validator, len(validates))
	for i := range validates {
		value := val.GetMapValue(validates[i].name)
		if value.IsNull() {
			if strict {
				val.err = errorf("no have [%s]  field to validate", validates[i].name)
				return validators{value: val}
			}
			continue
		}
		validates[i].value = value
		m[validates[i].name] = validates[i]
	}
	return validators{
		value:  val,
		values: m,
	}
}

func (val Value) Processors(ps ...processor) processors {
	m := make(map[string]processor, len(ps))
	for i := range ps {
		value := val.GetMapValue(ps[i].name)
		if value.IsNull() {
			if strict {
				val.err = errorf("no have [%s]  field to process", ps[i].name)
				return processors{value: val}
			}
			continue
		}
		ps[i].value = value
		m[ps[i].name] = ps[i]
	}
	return processors{
		value:  val,
		values: m,
	}
}

func (val Value) Processor(name string, apply ...Apply) processor {
	if err := val.GetError(); err != nil {
		return processor{value: val, name: name}
	}
	if val.IsMapValue() {
		val.err = errors.New("not support map value processor")
		return processor{value: val, name: name}
	}
	return processor{
		applies: apply,
		name:    name,
		value:   val,
	}

}

func (val Value) GetMapValue(name string) Value {
	if val.IsMapValue() {
		v, ok := val.v.(map[string]interface{})
		if ok {
			tt, ok2 := val.ty.typeImpl.(typeStringMap)
			if ok2 {
				return Value{
					ty: tt.GetStringMapType(name),
					v:  v[name],
				}
			}
		}
	}

	// todo
	return Value{
		err: errorf("%s no have value", name),
	}
}

func (val Value) SetMapValue(name string, value Value) Value {
	if val.IsMapValue() {
		v, ok := val.v.(map[string]interface{})
		if ok {
			_, ok2 := val.ty.typeImpl.(typeStringMap)
			if ok2 {
				v[name] = value.v
				val.ty.UpdateAttrType(map[string]Type{name: value.ty})
				return val
			}
		}
	}
	// todo
	return val
}

// 验证结果后直接赋值
func (val Value) Aligns(aligns ...align) error {
	if err := val.GetError(); err != nil {
		return err
	}
	if !val.IsMapValue() {
		return errors.New("single value is not support aligns")
	}
	for i := range aligns {
		if err := val.GetMapValue(aligns[i].name).Align(aligns[i]); err != nil {
			return err
		}
	}
	return nil
}

func (val Value) Align(a align) error {
	if err := val.GetError(); err != nil {
		return err
	}
	if err := val.UnMarshal(a.a); err != nil {
		return fmt.Errorf("align variable %s value error", a.name)
	}
	return nil
}

func (val Value) GetErrorResponseWriter(resp http.ResponseWriter) error {
	if val.err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(val.err.Error()))
		return val.err
	}
	return nil
}

func (val Value) String() string {
	return fmt.Sprint(val.v)
}

func (val Value) isNumber() bool {
	switch val.ty {
	case Int:
		return true
	case Int8:
		return true
	case Int16:
		return true
	case Int32:
		return true
	case Int64:
		return true
	case Uint:
		return true
	case Uint8:
		return true
	case Uint16:
		return true
	case Uint32:
		return true
	case Uint64:
		return true
	case Float64:
		return true
	case Float32:
		return true
	default:
		return false
	}
}
func (val Value) isString() bool {
	return val.ty == String
}
func (val Value) isBool() bool {
	return val.ty == Bool
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "value: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "value: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "value: Unmarshal(nil " + e.Type.String() + ")"
}
