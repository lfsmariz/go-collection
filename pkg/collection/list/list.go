package list

import (
	"github.com/go-collection/pkg/collection"
	"github.com/go-collection/pkg/element"
)

type List interface {
	collection.Collection
	AddAllAt(int, collection.Collection) bool
	ReplaceAll(func(element.Element) bool)
	Sort(func() int)
	Get(int) element.Element
	Set(int) element.Element
	AddTo(int, element.Element)
	RemoveFrom(int) element.Element
	IndexOf(element.Element) int
	LastIndexOf(element.Element) int
	SubList(int, int) List
	//TODO ListIterator
}
