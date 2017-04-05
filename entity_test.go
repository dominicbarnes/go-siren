package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	e := NewEntity("http://example.com")
	assert.Equal(t, "http://example.com", e.baseURL)
}

func TestEntityWithTitle(t *testing.T) {
	e := NewEntity("").WithTitle("hello world")
	assert.Equal(t, "hello world", e.Title)
}

func TestEntityWithTitleMultiple(t *testing.T) {
	e := NewEntity("").WithTitle("a").WithTitle("b")
	assert.Equal(t, "b", e.Title)
}

func TestEntityWithClasses(t *testing.T) {
	e := NewEntity("").WithClasses("a", "b", "c")
	assert.EqualValues(t, []string{"a", "b", "c"}, e.Class)
}

func TestEntityWithClassesMultiple(t *testing.T) {
	e := NewEntity("").WithClasses("a", "b").WithClasses("c", "d")
	assert.EqualValues(t, []string{"a", "b", "c", "d"}, e.Class)
}

func TestEntityWithProperties(t *testing.T) {
	e := NewEntity("").WithProperties(Properties{"a": "A"})
	assert.EqualValues(t, Properties{"a": "A"}, e.Properties)
}

func TestEntityWithPropertiesMultiple(t *testing.T) {
	e := NewEntity("").WithProperties(Properties{"a": "A"}).WithProperties(Properties{"b": "B"})
	assert.EqualValues(t, Properties{"a": "A", "b": "B"}, e.Properties)
}

func TestEntityWithProperty(t *testing.T) {
	e := NewEntity("").WithProperty("a", "A")
	assert.EqualValues(t, Properties{"a": "A"}, e.Properties)
}

func TestEntityWithPropertyMultiple(t *testing.T) {
	e := NewEntity("").WithProperty("a", "A").WithProperty("b", "B")
	assert.EqualValues(t, Properties{"a": "A", "b": "B"}, e.Properties)
}

func TestEntityWithLink(t *testing.T) {
	l := NewLink([]string{"self"}, "/")
	e := NewEntity("").WithLink(l)
	assert.EqualValues(t, []Link{*l}, e.Links)
}

func TestEntityWithLinkMultiple(t *testing.T) {
	l1 := NewLink([]string{"prev"}, "/posts/1")
	l2 := NewLink([]string{"next"}, "/posts/3")
	e := NewEntity("").WithLink(l1).WithLink(l2)
	assert.EqualValues(t, []Link{*l1, *l2}, e.Links)
}

func TestEntityWithAction(t *testing.T) {
	a := NewAction("create", "POST", "/posts")
	e := NewEntity("").WithAction(a)
	assert.EqualValues(t, []Action{*a}, e.Actions)
}

func TestEntityWithActionMultiple(t *testing.T) {
	a1 := NewAction("update", "PATCH", "/posts/1")
	a2 := NewAction("delete", "DELETE", "/posts/1")
	e := NewEntity("").WithAction(a1).WithAction(a2)
	assert.EqualValues(t, []Action{*a1, *a2}, e.Actions)
}

func TestEntityEmbed(t *testing.T) {
	ee := NewEmbeddedEntity([]string{"item"})
	e := NewEntity("").Embed(ee)
	assert.EqualValues(t, []EmbeddedEntity{*ee}, e.Entities)
}

func TestEntityValidate(t *testing.T) {
	e := NewEntity("")
	assert.NoError(t, e.Validate())
}

func TestEntityValidateEntities(t *testing.T) {
	ee := NewEmbeddedEntity(nil)
	e := NewEntity("").Embed(ee)
	assert.Error(t, e.Validate())
}

func TestEntityValidateLinks(t *testing.T) {
	e := NewEntity("").WithLink(NewLink(nil, "/"))
	assert.Error(t, e.Validate())
}

func TestEntityValidateActions(t *testing.T) {
	e := NewEntity("").WithAction(NewAction("", "GET", "/"))
	assert.Error(t, e.Validate())
}

func ExampleEntity() {
	NewEntity("http://api.x.io").
		WithClasses("order").
		WithProperties(Properties{
			"orderNumber": 42,
			"itemCount":   3,
			"status":      "pending",
		}).
		Embed(
			NewEmbeddedLink([]string{"http://x.io/rels/order-items"}, "/orders/42/items").
				WithClasses("items", "collection"),
		).
		Embed(
			NewEmbeddedEntity([]string{"http://x.io/rels/customer"}).
				WithProperties(Properties{
					"customerId": "pj123",
					"name":       "Peter Joseph",
				}).
				WithLink(NewLink([]string{"self"}, "/customers/pj123")).
				WithClasses("info", "customer"),
		).
		WithAction(NewAction("add-item", "POST", "/orders/42/items").
			WithTitle("Add Item").
			WithType("application/x-www-form-urlencoded").
			WithField(NewActionField("orderNumber", "hidden").WithValue("42")).
			WithField(NewActionField("productCode", "text")).
			WithField(NewActionField("quantity", "number"))).
		WithLink(NewLink([]string{"self"}, "/orders/42")).
		WithLink(NewLink([]string{"previous"}, "/orders/41")).
		WithLink(NewLink([]string{"next"}, "/orders/43"))
}
