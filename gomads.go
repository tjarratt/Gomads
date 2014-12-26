package gomads

type maybeValue struct {
	v interface{}
}

func (m maybeValue) Value() interface{} {
	return m.v
}

func (m maybeValue) OrSome(v interface{}) MaybeValue {
	if m.v == nil {
		m.v = v
	}

	return m
}

type MaybeValue interface {
	Value() interface{}
	OrSome(interface{}) MaybeValue
}

func Maybe(unit func() interface{}) MaybeValue {
	return maybeValue{v: unit()}
}
