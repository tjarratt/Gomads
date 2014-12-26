package gomads

type maybeValue struct {
	v interface{}
}

func (m maybeValue) Value() interface{} {
	return m.v
}

type MaybeValue interface {
	Value() interface{}
}

func Maybe(unit func() interface{}) MaybeValue {
	return maybeValue{v: unit()}
}
