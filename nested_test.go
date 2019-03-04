package nested

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Imagine we have slices inside structs inside slices and structs.
// Wanted just easy example here.
type Nested struct {
	Foo []struct {
		Bar int
	}
}

func TestNestedDots_Empty(t *testing.T) {
	s := &Nested{}
	value := Resolve(666, func() interface{} {
		return s.Foo[0].Bar
	})
	assert.Equal(t, 666, value)
}

func TestNestedDots_Filled(t *testing.T) {
	s := &Nested{
		Foo: []struct{ Bar int }{{Bar: 5}},
	}
	value := Resolve(666, func() interface{} {
		return s.Foo[0].Bar
	})
	assert.Equal(t, 5, value)
}
