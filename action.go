package siren

import (
	"net/http"

	validator "gopkg.in/validator.v2"
)

const (
	// ActionDefaultMethod is the default method for an action when one is not
	// specified by the server.
	ActionDefaultMethod = http.MethodGet

	// ActionDefaultType is the default media type for the action when one is not
	// specified by the server.
	ActionDefaultType = "application/x-www-form-urlencoded"
)

// Action is a description of an action to take on a related resource.
type Action struct {
	Name   string        `json:"name" validate:"nonzero"`
	Href   Href          `json:"href" validate:"nonzero"`
	Method string        `json:"method,omitempty"`
	Fields []ActionField `json:"fields,omitempty"`
	Type   string        `json:"type,omitempty"`
	Title  string        `json:"title,omitempty"`
	Class  Classes       `json:"class,omitempty"`
}

// Validate ensures that the link is well-formed.
func (a Action) Validate() error {
	return validator.Validate(a)
}

// GetMethod is a helper for getting the HTTP method for an action. When the
// action is not explicit, ActionDefaultMethod will be returned. This is a
// convenience method offered for clients.
func (a Action) GetMethod() string {
	if a.Method == "" {
		return ActionDefaultMethod
	}

	return a.Method
}

// GetType is a helper for getting the media type for the action request body.
// When the action is not explicit, ActionDefaultType will be returned. This is
// a convenience method offered for clients.
func (a Action) GetType() string {
	if a.Type == "" {
		return ActionDefaultType
	}

	return a.Type
}

// WithBaseHref returns a copy of this action that applies the supplied base
// href when the action's href starts with a /.
func (a Action) WithBaseHref(base Href) Action {
	return Action{
		Name:   a.Name,
		Href:   a.Href.WithBaseHref(base),
		Method: a.Method,
		Fields: a.Fields,
		Type:   a.Type,
		Title:  a.Title,
		Class:  a.Class,
	}
}

// ActionField is a single field within the larger action.
type ActionField struct {
	Name  string  `json:"name" validate:"nonzero"`
	Type  string  `json:"type,omitempty"`
	Value any     `json:"value,omitempty"`
	Title string  `json:"title,omitempty"`
	Class Classes `json:"class,omitempty"`
}
