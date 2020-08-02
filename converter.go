package optional

import (
	"encoding/json"
	"errors"
	"strconv"
)

type converter struct {
	typeConverter
	value Value
}

type typeConverter interface {
	Int() (int, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)
	String() (string, error)
	Float32() (float32, error)
	Float64() (float64, error)
	Bool() (bool, error)
	Belief
}

var ConvertError = errors.New("This type conversion is not supported")

func (c converter) Int() (int, error) {
	switch c.value.ty {
	case Int:
		i, ok := c.value.v.(int)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	case Int8:
		i, ok := c.value.v.(int8)
		if ok {
			return int(i), nil
		}
		return 0, ConvertError
	case Int16:
		i, ok := c.value.v.(int16)
		if ok {
			return int(i), nil
		}
		return 0, ConvertError
	case Int32:
		i, ok := c.value.v.(int32)
		if ok {
			return int(i), nil
		}
		return 0, ConvertError
	case Int64:
		i, ok := c.value.v.(int64)
		if ok {
			return int(i), nil
		}
		return 0, ConvertError
	case String:
		return strconv.Atoi(c.value.v.(string))
	default:
		return 0, ConvertError
	}
}
func (c converter) Int8() (int8, error) {
	if c.value.ty == Int8 {
		i, ok := c.value.v.(int8)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	i, err := c.Int()
	return int8(i), err
}
func (c converter) Int16() (int16, error) {
	if c.value.ty == Int16 {
		i, ok := c.value.v.(int16)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	i, err := c.Int()
	return int16(i), err
}
func (c converter) Int32() (int32, error) {
	if c.value.ty == Int32 {
		i, ok := c.value.v.(int32)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	if c.value.ty == Int64 {
		i, ok := c.value.v.(int64)
		if ok {
			return int32(i), nil
		}
		return 0, ConvertError
	}
	if c.value.ty == String {
		i64, err := strconv.ParseInt(c.value.v.(string), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i64), nil
	}
	i, err := c.Int()
	return int32(i), err
}
func (c converter) Int64() (int64, error) {
	if c.value.ty == Int64 {
		i, ok := c.value.v.(int64)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	if c.value.ty == String {
		return strconv.ParseInt(c.value.v.(string), 10, 64)
	}
	i, err := c.Int()
	return int64(i), err
}

func (c converter) Uint() (uint, error) {
	switch c.value.ty {
	case Uint:
		i, ok := c.value.v.(uint)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	case Uint8:
		i, ok := c.value.v.(uint8)
		if ok {
			return uint(i), nil
		}
		return 0, ConvertError
	case Uint16:
		i, ok := c.value.v.(uint16)
		if ok {
			return uint(i), nil
		}
		return 0, ConvertError
	case Uint32:
		i, ok := c.value.v.(uint32)
		if ok {
			return uint(i), nil
		}
		return 0, ConvertError
	case Uint64:
		i, ok := c.value.v.(uint64)
		if ok {
			return uint(i), nil
		}
		return 0, ConvertError
	default:
		return 0, ConvertError
	}
}
func (c converter) Uint8() (uint8, error) {
	if c.value.ty == Uint8 {
		i, ok := c.value.v.(uint8)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	i, err := c.Uint()
	return uint8(i), err
}
func (c converter) Uint16() (uint16, error) {
	if c.value.ty == Uint16 {
		i, ok := c.value.v.(uint16)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	i, err := c.Uint()
	return uint16(i), err
}
func (c converter) Uint32() (uint32, error) {
	if c.value.ty == Uint32 {
		i, ok := c.value.v.(uint32)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}

	i, err := c.Uint()
	return uint32(i), err
}
func (c converter) Uint64() (uint64, error) {
	if c.value.ty == Uint64 {
		i, ok := c.value.v.(uint64)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	i, err := c.Uint()
	return uint64(i), err
}

func (c converter) Float32() (float32, error) {
	switch c.value.ty {
	case Float32:
		i, ok := c.value.v.(float32)
		if ok {
			return i, nil
		}
	case Float64:
		i, ok := c.value.v.(float64)
		if ok {
			return float32(i), nil
		}
	case String:
		f64, err := strconv.ParseFloat(c.value.v.(string), 64)
		if err != nil {
			return 0, err
		}
		return float32(f64), err
	}
	return 0, ConvertError
}
func (c converter) Float64() (float64, error) {
	if c.value.ty == Float64 {
		i, ok := c.value.v.(float64)
		if ok {
			return i, nil
		}
		return 0, ConvertError
	}
	if c.value.ty == String {
		return strconv.ParseFloat(c.value.v.(string), 64)
	}
	i, err := c.Float32()

	return float64(i), err
}

func (c converter) isNumber() bool {
	return c.value.isNumber()
}
func (c converter) isBool() bool {
	return c.value.ty == Bool
}
func (c converter) isString() bool {
	return c.value.ty == String
}

func (c converter) String() (string, error) {
	if c.value.ty.IsPrimitiveType() {
		if c.value.ty == String {
			return c.value.v.(string), nil
		}
		if c.value.ty == Bool {
			return strconv.FormatBool(c.value.v.(bool)), nil
		}
		if c.isNumber() {
			if c.value.ty == Float32 {
				f64, err := c.Float64()
				if err != nil {
					return "", err
				}
				return strconv.FormatFloat(f64, 'E', -1, 32), nil
			}
			if c.value.ty == Float64 {
				f64, err := c.Float64()
				if err != nil {
					return "", err
				}
				return strconv.FormatFloat(f64, 'E', -1, 64), nil
			}

			i64, err := c.Int64()
			if err != nil {
				return "", err
			}
			return strconv.FormatInt(i64, 10), nil
		}
	}
	if c.value.ty.IsMapType() {
		//todo
		v, err := json.Marshal(c.value.v)
		return string(v), err
	}
	return "", ConvertError
}
func (c converter) Bool() (bool, error) {
	switch c.value.ty {
	case Bool:
		b, ok := c.value.v.(bool)
		if ok {
			return b, nil
		}
		return false, ConvertError
	case String:
		return strconv.ParseBool(c.value.v.(string))
	default:
		return false, ConvertError
	}
}

func (c converter) JSON() ([]byte, error) {
	if c.value.ty.IsMapType() {
		return json.Marshal(c.value.v)
	}
	return nil, ConvertError
}

func (c converter) Value() Value {
	return c.value
}

func (c converter) GetError() error {
	return c.value.err
}
