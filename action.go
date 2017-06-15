package siren

import "gopkg.in/validator.v2"

// Action is a description of an action to take on a related resource.
type Action struct {
	Name   string        `json:"name" validate:"nonzero"`
	Href   string        `json:"href" validate:"nonzero"`
	Method string        `json:"method,omitempty"`
	Fields []ActionField `json:"fields,omitempty"`
	Type   string        `json:"type,omitempty"`
	Title  string        `json:"title,omitempty"`
	Class  Classes       `json:"class,omitempty"`
}

// NewAction is a helper for creating a new Action instance.
func NewAction(name, method, href string) *Action {
	return &Action{
		Name:   name,
		Method: method,
		Href:   href,
	}
}

// WithType is a helper for setting an optional media type for this resource action.
// Calling multiple times will replace any previous values.
func (a *Action) WithType(mediaType string) *Action {
	a.Type = mediaType
	return a
}

// WithTitle is a helper for setting an optional title for this resource action.
// Calling multiple times will replace any previous values.
func (a *Action) WithTitle(title string) *Action {
	a.Title = title
	return a
}

// WithClasses appends new class names to the resource action.
// Calling multiple times will continuously append.
func (a *Action) WithClasses(classes ...string) *Action {
	a.Class = append(a.Class, classes...)
	return a
}

// WithField appends a new field to the resource action.
// Calling multiple times will continuously append.
func (a *Action) WithField(field *ActionField) *Action {
	a.Fields = append(a.Fields, *field)
	return a
}

// Validate ensures the action is well-formed.
func (a *Action) Validate() error {
	return validator.Validate(a)
}

// ActionField is a single field within the larger action.
type ActionField struct {
	Name  string      `json:"name" validate:"nonzero"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Title string      `json:"title,omitempty"`
	Class Classes     `json:"class,omitempty"`
}

// NewActionField is a helper for creating a new ActionField instance.
func NewActionField(name, fieldType string) *ActionField {
	return &ActionField{
		Name: name,
		Type: fieldType,
	}
}

// WithTitle is a helper for setting an optional title for this action field.
// Calling multiple times will replace any previous values.
func (f *ActionField) WithTitle(title string) *ActionField {
	f.Title = title
	return f
}

// WithClasses appends new class names to the action field.
// Calling multiple times will continuously append.
func (f *ActionField) WithClasses(classes ...string) *ActionField {
	f.Class = append(f.Class, classes...)
	return f
}

// WithValue sets the value for the action field.
// Calling multiple times will replace any previous value. (nil is acceptable)
func (f *ActionField) WithValue(value interface{}) *ActionField {
	f.Value = value
	return f
}

// Validate ensures the action field is well-formed.
func (f *ActionField) Validate() error {
	return validator.Validate(f)
}
