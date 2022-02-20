package iterator

import "github.com/go-collection/pkg/element"

type Iterator interface {
	HasNext() bool

	Next() element.Element

	Remove()

	ForEachRemaining(action func(element.Element))
}
