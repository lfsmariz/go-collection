package iterable

import (
	"github.com/go-collection/internal/element"
	"github.com/go-collection/internal/iterator"
)

type Iterable interface {
	Iterator() iterator.Iterator

	ForEach(action func(element.Element))
}
