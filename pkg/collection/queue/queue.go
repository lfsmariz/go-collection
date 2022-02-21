package queue

import (
	"github.com/go-collection/pkg/collection"
	"github.com/go-collection/pkg/element"
)

type Queue interface {
	collection.Collection
	Offer(element.Element) bool
	RemoveQueue() (element.Element, error)
	Poll() element.Element
	First() element.Element
	Peek() element.Element
}
