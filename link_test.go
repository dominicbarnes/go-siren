package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	l := NewLink(Rels{"self"}, "/")
	assert.EqualValues(t, Rels{"self"}, l.Rel)
	assert.Equal(t, "/", l.Href)
}

func TestLinkWithTitle(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithTitle("hello world")
	assert.Equal(t, "hello world", l.Title)
}

func TestLinkWithTitleMultiple(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithTitle("a").WithTitle("b")
	assert.Equal(t, "b", l.Title)
}

func TestLinkWithType(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithType("text/plain")
	assert.Equal(t, "text/plain", l.Type)
}

func TestLinkWithTypeMultiple(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithType("a").WithType("b")
	assert.Equal(t, "b", l.Type)
}

func TestLinkWithClasses(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithClasses("a", "b", "c")
	assert.EqualValues(t, Rels{"a", "b", "c"}, l.Class)
}

func TestLinkWithClassesMultiple(t *testing.T) {
	l := NewLink(Rels{"self"}, "/").WithClasses("a", "b").WithClasses("c", "d")
	assert.EqualValues(t, Rels{"a", "b", "c", "d"}, l.Class)
}

func TestLinkValidate(t *testing.T) {
	l := NewLink(Rels{"self"}, "/")
	assert.NoError(t, l.Validate())
}

func TestLinkValidateRel(t *testing.T) {
	l := NewLink(Rels{}, "/")
	assert.EqualError(t, l.Validate(), "Rel: zero value")
}

func TestLinkValidateHref(t *testing.T) {
	l := NewLink(Rels{"self"}, "")
	assert.EqualError(t, l.Validate(), "Href: zero value")
}
