package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertiesMerge(t *testing.T) {
	p := Properties{"a": 1}
	p.Merge(Properties{"b": 2})

	assert.EqualValues(t, Properties{"a": 1, "b": 2}, p)
}

func TestPropertiesMergeOverwrite(t *testing.T) {
	p := Properties{"a": 1, "b": 2}
	p.Merge(Properties{"a": 3})

	assert.EqualValues(t, Properties{"a": 3, "b": 2}, p)
}
