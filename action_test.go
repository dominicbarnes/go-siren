package siren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Action

func TestNewAction(t *testing.T) {
	a := NewAction("create", "POST", "/posts")
	assert.Equal(t, "create", a.Name)
	assert.Equal(t, "POST", a.Method)
	assert.Equal(t, "/posts", a.Href)
}

func TestActionWithTitle(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithTitle("hello world")
	assert.Equal(t, "hello world", a.Title)
}

func TestActionWithTitleMultiple(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithTitle("a").WithTitle("b")
	assert.Equal(t, "b", a.Title)
}

func TestActionWithType(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithType("plain/text")
	assert.Equal(t, "plain/text", a.Type)
}

func TestActionWithTypeMultiple(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithType("a").WithType("b")
	assert.Equal(t, "b", a.Type)
}

func TestActionWithClasses(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithClasses("a", "b", "c")
	assert.EqualValues(t, []string{"a", "b", "c"}, a.Class)
}

func TestActionWithClassesMultiple(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithClasses("a", "b").WithClasses("c", "d")
	assert.EqualValues(t, []string{"a", "b", "c", "d"}, a.Class)
}

func TestActionWithField(t *testing.T) {
	a := NewAction("create", "POST", "/posts").WithField(NewActionField("name", "text"))
	assert.EqualValues(t, []ActionField{{Name: "name", Type: "text"}}, a.Fields)
}

func TestActionValidate(t *testing.T) {
	a := NewAction("create", "POST", "/posts")
	assert.NoError(t, a.Validate())
}

func TestActionValidateName(t *testing.T) {
	a := NewAction("", "POST", "/posts")
	assert.Error(t, a.Validate())
}

func TestActionValidateHref(t *testing.T) {
	a := NewAction("create", "POST", "")
	assert.Error(t, a.Validate())
}

// ActionField

func TestNewActionField(t *testing.T) {
	f := NewActionField("name", "text")
	assert.Equal(t, "name", f.Name)
	assert.Equal(t, "text", f.Type)
}

func TestActionFieldWithTitle(t *testing.T) {
	f := NewActionField("name", "text").WithTitle("Name")
	assert.Equal(t, "Name", f.Title)
}

func TestActionFieldWithTitleMultiple(t *testing.T) {
	f := NewActionField("name", "text").WithTitle("A").WithTitle("B")
	assert.Equal(t, "B", f.Title)
}

func TestActionFieldWithClasses(t *testing.T) {
	f := NewActionField("name", "text").WithClasses("a", "b", "c")
	assert.EqualValues(t, []string{"a", "b", "c"}, f.Class)
}

func TestActionFieldWithClassesMultiple(t *testing.T) {
	f := NewActionField("name", "text").WithClasses("a", "b").WithClasses("c", "d")
	assert.EqualValues(t, []string{"a", "b", "c", "d"}, f.Class)
}

func TestActionFieldWithValue(t *testing.T) {
	f := NewActionField("name", "text").WithValue("John Doe")
	assert.Equal(t, "John Doe", f.Value)
}

func TestActionFieldWithValueMultiple(t *testing.T) {
	f := NewActionField("name", "text").WithValue("A").WithValue("B")
	assert.Equal(t, "B", f.Value)
}

func TestActionFieldValidate(t *testing.T) {
	f := NewActionField("name", "text")
	assert.NoError(t, f.Validate())
}

func TestActionFieldValidateName(t *testing.T) {
	f := NewActionField("", "text")
	assert.Error(t, f.Validate())
}
