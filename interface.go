package optional


type Returner interface {
	HashError() bool
	GetError() error
}

type Converter interface {
}

type Validator interface {
	IsString() bool
}

type Processor interface {
}
type Aligner interface {
}

