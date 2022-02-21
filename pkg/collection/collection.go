package collection

import (
	iterable "github.com/go-collection/internal/Iterable"
	"github.com/go-collection/pkg/element"
)

type Collection interface {
	iterable.Iterable
	Size() int
	IsEmpty() bool
	Contains(element.Element) bool
	ToArray() []element.Element
	Add(element.Element) bool
	Remove(element.Element) bool
	ContainsAll(Collection) bool
	AddAll(Collection) bool
	RemoveAll(Collection) bool
	RemoveIf(func(element.Element) bool)
	RetainAll(Collection) bool
	Clear()
	Equals() bool
}
