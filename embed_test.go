package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbeddedEntity(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"})
	assert.EqualValues(t, []string{"item"}, e.Rel)
}

func TestEmbeddedLink(t *testing.T) {
	e := NewEmbeddedLink([]string{"item"}, "/posts/1")
	assert.EqualValues(t, []string{"item"}, e.Rel)
	assert.EqualValues(t, "/posts/1", e.Href)
}

func TestEmbeddedEntityWithTitle(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithTitle("Item")
	assert.Equal(t, "Item", e.Title)
}

func TestEmbeddedEntityWithProperties(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithProperties(Properties{"a": 1})
	assert.Equal(t, Properties{"a": 1}, e.Properties)
}

func TestEmbeddedEntityWithProperty(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithProperty("a", 1)
	assert.Equal(t, Properties{"a": 1}, e.Properties)
}

func TestEmbeddedEntityWithClasses(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithClasses("a", "b")
	assert.Equal(t, []string{"a", "b"}, e.Class)
}

func TestEmbeddedEntityWithLink(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithLink(NewLink([]string{"self"}, "/posts/1"))
	assert.Len(t, e.Links, 1)
}

func TestEmbeddedEntityWithAction(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"}).WithAction(NewAction("create", "POST", "/posts"))
	assert.Len(t, e.Actions, 1)
}

func TestEmbeddedEntityValidate(t *testing.T) {
	e := NewEmbeddedEntity([]string{"item"})
	assert.NoError(t, e.Validate())
}

func TestEmbeddedEntityValidateRel(t *testing.T) {
	e := NewEmbeddedEntity(nil)
	assert.Error(t, e.Validate())
}

func TestEmbeddedLinkValidateRel(t *testing.T) {
	e := NewEmbeddedLink(nil, "")
	assert.Error(t, e.Validate())
}
