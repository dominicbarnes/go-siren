package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	l := NewLink([]string{"self"}, "/")
	assert.EqualValues(t, []string{"self"}, l.Rel)
	assert.Equal(t, "/", l.Href)
}

func TestLinkWithTitle(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithTitle("hello world")
	assert.Equal(t, "hello world", l.Title)
}

func TestLinkWithTitleMultiple(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithTitle("a").WithTitle("b")
	assert.Equal(t, "b", l.Title)
}

func TestLinkWithType(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithType("text/plain")
	assert.Equal(t, "text/plain", l.Type)
}

func TestLinkWithTypeMultiple(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithType("a").WithType("b")
	assert.Equal(t, "b", l.Type)
}

func TestLinkWithClasses(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithClasses("a", "b", "c")
	assert.EqualValues(t, []string{"a", "b", "c"}, l.Class)
}

func TestLinkWithClassesMultiple(t *testing.T) {
	l := NewLink([]string{"self"}, "/").WithClasses("a", "b").WithClasses("c", "d")
	assert.EqualValues(t, []string{"a", "b", "c", "d"}, l.Class)
}

func TestLinkValidate(t *testing.T) {
	l := NewLink([]string{"self"}, "/")
	assert.NoError(t, l.Validate())
}

func TestLinkValidateRel(t *testing.T) {
	l := NewLink([]string{}, "/")
	assert.EqualError(t, l.Validate(), "Rel: zero value")
}

func TestLinkValidateHref(t *testing.T) {
	l := NewLink([]string{"self"}, "")
	assert.EqualError(t, l.Validate(), "Href: zero value")
}
