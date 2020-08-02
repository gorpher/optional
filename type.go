package optional

// Type represents value types within the type system.
//
// This is a closed interface type, meaning that only the concrete
// implementations provided within this package are considered valid.
type Type struct {
	typeImpl
}

type typeImpl interface {
	// isTypeImpl is a do-nothing method that exists only to express
	// that a type is an implementation of typeImpl.
	isTypeImpl() typeImplSigil

	// Equals returns true if the other given Type exactly equals the
	// receiver Type.
	Equals(other Type) bool

	// FriendlyName returns a human-friendly *English* name for the given
	// type.
	FriendlyName() string

	// GoString implements the GoStringer interface from package fmt.
	GoString() string
}

type typeImplSigil struct{}

func (t typeImplSigil) isTypeImpl() typeImplSigil {
	return typeImplSigil{}
}

func (t Type) Equals(other Type) bool {
	return t.typeImpl.Equals(other)
}

func (t Type) GoString() string {
	if t.typeImpl == nil {
		return "optional.NilType"
	}

	return t.typeImpl.GoString()
}

func (t Type) FriendlyName() string {
	return t.typeImpl.FriendlyName()
}
