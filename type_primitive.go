package optional

import "math/big"

type primitiveType struct {
	typeImplSigil
	Kind primitiveTypeKind
}

type primitiveTypeKind string

const (
	primitiveTypeBool    primitiveTypeKind = "bool"
	primitiveTypeString  primitiveTypeKind = "string"
	primitiveTypeInt     primitiveTypeKind = "int"
	primitiveTypeUint    primitiveTypeKind = "uint"
	primitiveTypeUint8   primitiveTypeKind = "uint8"
	primitiveTypeUint16  primitiveTypeKind = "uint16"
	primitiveTypeUint32  primitiveTypeKind = "uint32"
	primitiveTypeUint64  primitiveTypeKind = "uint64"
	primitiveTypeInt8    primitiveTypeKind = "int8"
	primitiveTypeInt16   primitiveTypeKind = "int16"
	primitiveTypeInt32   primitiveTypeKind = "int32"
	primitiveTypeInt64   primitiveTypeKind = "int64"
	primitiveTypeFloat32 primitiveTypeKind = "float32"
	primitiveTypeFloat64 primitiveTypeKind = "float64"
)

func (t primitiveType) Equals(other Type) bool {
	if otherP, ok := other.typeImpl.(primitiveType); ok {
		return otherP.Kind == t.Kind
	}
	return false
}
func (t primitiveType) FriendlyName() string {
	switch t.Kind {
	case primitiveTypeBool:
		return string(primitiveTypeBool)
	case primitiveTypeInt:
		return string(primitiveTypeInt)
	case primitiveTypeUint:
		return string(primitiveTypeUint)
	case primitiveTypeInt8:
		return string(primitiveTypeInt8)
	case primitiveTypeUint8:
		return string(primitiveTypeUint8)
	case primitiveTypeInt16:
		return string(primitiveTypeInt16)
	case primitiveTypeUint16:
		return string(primitiveTypeUint16)
	case primitiveTypeInt32:
		return string(primitiveTypeInt32)
	case primitiveTypeUint32:
		return string(primitiveTypeUint32)
	case primitiveTypeInt64:
		return string(primitiveTypeInt64)
	case primitiveTypeUint64:
		return string(primitiveTypeUint64)
	case primitiveTypeFloat32:
		return string(primitiveTypeFloat32)
	case primitiveTypeFloat64:
		return string(primitiveTypeFloat64)
	case primitiveTypeString:
		return string(primitiveTypeString)
	default:
		// should never happen
		panic("invalid primitive type")
	}
}

func (t primitiveType) GoString() string {
	switch t.Kind {
	case primitiveTypeBool:
		return "optional.Bool"
	case primitiveTypeInt:
		return "optional.Int"
	case primitiveTypeInt8:
		return "optional.Int8"
	case primitiveTypeInt16:
		return "optional.Int16"
	case primitiveTypeInt32:
		return "optional.Int32"
	case primitiveTypeInt64:
		return "optional.Int64"
	case primitiveTypeUint:
		return "optional.Uint"
	case primitiveTypeUint8:
		return "optional.Uint8"
	case primitiveTypeUint16:
		return "optional.Uint16"
	case primitiveTypeUint32:
		return "optional.Uint32"
	case primitiveTypeUint64:
		return "optional.Uint64"
	case primitiveTypeString:
		return "optional.String"
	case primitiveTypeFloat32:
		return "optional.Float32"
	case primitiveTypeFloat64:
		return "optional.Float64"

	default:
		// should never happen
		panic("invalid primitive type")
	}
}

var Int Type
var Uint Type
var Uint8 Type
var Uint16 Type
var Uint32 Type
var Uint64 Type
var Int8 Type
var Int16 Type
var Int32 Type
var Int64 Type
var Float32 Type
var Float64 Type

var String Type

var Bool Type

var True Value

var False Value

var Zero Value

var PositiveInfinity Value

var NegativeInfinity Value

func init() {
	Int = Type{
		primitiveType{Kind: primitiveTypeInt},
	}
	Uint = Type{
		primitiveType{Kind: primitiveTypeUint},
	}
	Uint8 = Type{
		primitiveType{Kind: primitiveTypeUint8},
	}
	Uint16 = Type{
		primitiveType{Kind: primitiveTypeUint16},
	}
	Uint32 = Type{
		primitiveType{Kind: primitiveTypeUint32},
	}
	Uint64 = Type{
		primitiveType{Kind: primitiveTypeUint64},
	}
	Int8 = Type{
		primitiveType{Kind: primitiveTypeInt8},
	}
	Int16 = Type{
		primitiveType{Kind: primitiveTypeInt16},
	}
	Int32 = Type{
		primitiveType{Kind: primitiveTypeInt32},
	}
	Int64 = Type{
		primitiveType{Kind: primitiveTypeInt64},
	}
	Float32 = Type{
		primitiveType{Kind: primitiveTypeFloat32},
	}
	Float64 = Type{
		primitiveType{Kind: primitiveTypeFloat64},
	}
	String = Type{
		primitiveType{Kind: primitiveTypeString},
	}
	Bool = Type{
		primitiveType{Kind: primitiveTypeBool},
	}

	True = Value{
		ty: Bool,
		v:  true,
	}
	False = Value{
		ty: Bool,
		v:  false,
	}
	Zero = Value{
		ty: Int,
		v:  0,
	}
	PositiveInfinity = Value{
		ty: Float64,
		v:  (&big.Float{}).SetInf(false),
	}
	NegativeInfinity = Value{
		ty: Float64,
		v:  (&big.Float{}).SetInf(true),
	}
}

func (t Type) IsPrimitiveType() bool {
	_, ok := t.typeImpl.(primitiveType)
	return ok
}
