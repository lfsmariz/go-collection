package iterator

import (
	"github.com/go-collection/internal/element"
)

type Iterator interface {
	HasNext() bool

	Next() element.Element

	Remove()

	ForEachRemaining(action func(element.Element))
}
