package sortedset

import (
	"github.com/go-collection/pkg/collection/set"
	"github.com/go-collection/pkg/element"
)

type SortedSet interface {
	set.Set
	Comparator() func()
	First() element.Element
	Last() element.Element
	HeadSet(element.Element) SortedSet
	SubSet(element.Element, element.Element) SortedSet
	TailSet(element.Element) SortedSet
}
