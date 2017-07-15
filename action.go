package siren

import (
	"net/http"

	validator "gopkg.in/validator.v2"
)

const (
	ActionDefaultMethod = http.MethodGet
	ActionDefaultType   = "application/x-www-form-urlencoded"
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

func (a Action) GetMethod() string {
	if a.Method == "" {
		return ActionDefaultMethod
	}

	return a.Method
}

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
	Name  string      `json:"name" validate:"nonzero"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Title string      `json:"title,omitempty"`
	Class Classes     `json:"class,omitempty"`
}
