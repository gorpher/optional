package optional

import (
	"errors"
	"fmt"
	"strconv"
)

// BoolVal returns a Value of type Number whose internal value is the given
// bool.
func BoolVal(v bool) Value {
	return Value{
		ty: Bool,
		v:  v,
	}
}

// ParseNumberVal returns a Value of type number produced by parsing the given
// string as a decimal real number. To ensure that two identical strings will
// always produce an equal number, always use this function to derive a number
// from a string; it will ensure that the precision and rounding mode for the
// internal big decimal is configured in a consistent way.
//
// If the given string cannot be parsed as a number, the returned error has
// the message "a number is required", making it suitable to return to an
// end-user to signal a type conversion error.
//
// If the given string contains a number that becomes a recurring fraction
// when expressed in binary then it will be truncated to have a 512-bit
// mantissa. Note that this is a higher precision than that of a float64,
// so coverting the same decimal number first to float64 and then calling
// NumberFloatVal will not produce an equal result; the conversion first
// to float64 will round the mantissa to fewer than 512 bits.
func ParseFloat64Val(s string) (Value, error) {
	int64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return NilVal, fmt.Errorf("a number is required")
	}
	return Value{ty: Float64, v: int64}, nil
}

func IntVal(v int) Value {
	return Value{
		ty: Int,
		v:  v,
	}
}
func Int8Val(v int8) Value {
	return Value{
		ty: Int8,
		v:  v,
	}
}
func Int16Val(v int16) Value {
	return Value{
		ty: Int16,
		v:  v,
	}
}
func Int32Val(v int32) Value {
	return Value{
		ty: Int32,
		v:  v,
	}
}
func Int64Val(v int64) Value {
	return Value{
		ty: Int64,
		v:  v,
	}
}
func UintVal(v uint) Value {
	return Value{
		ty: Uint,
		v:  v,
	}
}
func Uint8Val(v uint8) Value {
	return Value{
		ty: Uint8,
		v:  v,
	}
}
func Uint16Val(v uint16) Value {
	return Value{
		ty: Uint16,
		v:  v,
	}
}
func Uint32Val(v uint32) Value {
	return Value{
		ty: Uint32,
		v:  v,
	}
}
func Uint64Val(v uint64) Value {
	return Value{
		ty: Uint64,
		v:  v,
	}
}
func Float32Val(v float32) Value {
	return Value{
		ty: Float32,
		v:  v,
	}
}
func Float64Val(v float64) Value {
	return Value{
		ty: Float64,
		v:  v,
	}
}

// NullVal returns a null value of the given type. A null can be created of any
// type, but operations on such values will always panic. Calling applications
// are encouraged to use nulls only sparingly, particularly when user-provided
// expressions are to be evaluated, since the precence of nulls creates a
// much higher chance of evaluation errors that can't be caught by a type
// checker.
func NullVal(t Type) Value {
	return Value{
		ty: t,
		v:  nil,
	}
}

// StringVal returns a Value of type String whose internal value is the
// given string.
//
// Strings must be UTF-8 encoded sequences of valid unicode codepoints, and
// they are NFC-normalized on entry into the world of cty values.
//
// If the given string is not valid UTF-8 then behavior of string operations
// is undefined.
func StringVal(v string) Value {
	return Value{
		ty: String,
		v:  v,
	}
}

func MapStringVal(vals map[string]Value) Value {
	if len(vals) == 0 {
		return Value{err: errors.New("must not call MapVal with empty map")}
	}
	rawMap := make(map[string]interface{}, len(vals))
	rawType := make(map[string]Type, len(vals))
	for key, val := range vals {
		rawMap[key] = val.v
		rawType[key] = val.ty
	}
	return Value{
		ty: StringMapType(rawType),
		v:  rawMap,
	}
}

func MapStringValEmpty() Value {
	return Value{
		ty: StringMap(),
		v:  map[string]interface{}{},
	}
}
