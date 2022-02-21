package set

import (
	"github.com/go-collection/pkg/collection"
	"github.com/go-collection/pkg/element"
)

type Set interface {
	collection.Collection
	Of(...element.Element)
}
