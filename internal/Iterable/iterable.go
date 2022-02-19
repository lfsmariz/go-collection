package iterable

import (
	"github.com/go-collection/internal/iterator"
	"github.com/go-collection/pkg/element"
)

type Iterable interface {
	Iterator() iterator.Iterator

	ForEach(action func(element.Element))
}
