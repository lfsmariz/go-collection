package element

type Element interface {
	Value() Element
}

type ElementImpl struct {
	A int
}

func (e *ElementImpl) Value() Element {
	return e
}
