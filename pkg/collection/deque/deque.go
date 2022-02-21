package deque

import (
	"github.com/go-collection/pkg/collection/queue"
	"github.com/go-collection/pkg/element"
)

type Deque interface {
	queue.Queue
	AddFirst(element.Element)
	AddLast(element.Element)
	OfferFirst(element.Element) bool
	OfferLast(element.Element) bool
	RemoveFirst() element.Element
	RemoveLast() element.Element
	PollFirst() element.Element
	PollLast() element.Element
	GetFirst() element.Element
	GetLast() element.Element
	PeekFirst() element.Element
	PeekLast() element.Element
	RemoveFirstOccurrence(element.Element) bool
	RemoveLastOccurrence(element.Element) bool
	Push(element.Element)
	Pop(element.Element)
}
