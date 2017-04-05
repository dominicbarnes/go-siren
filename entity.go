package siren

import "gopkg.in/validator.v2"

// Entity is a top-level resource in a siren API.
type Entity struct {
	Entities   []EmbeddedEntity `json:"entities,omitempty"`
	Links      []Link           `json:"links,omitempty"`
	Actions    []Action         `json:"actions,omitempty"`
	Properties Properties       `json:"properties,omitempty"`
	Title      string           `json:"title,omitempty"`
	Class      []string         `json:"class,omitempty"`
	baseURL    string
}

// NewEntity is a helper for creating a new entity instance.
func NewEntity(base string) *Entity {
	return &Entity{baseURL: base}
}

// WithTitle will set the title for this entity.
// Calling multiple times will simply overwrite any previous value.
func (e *Entity) WithTitle(title string) *Entity {
	e.Title = title
	return e
}

// WithClasses appends new class names to the entity.
// Calling multiple times will continuously append.
func (e *Entity) WithClasses(classes ...string) *Entity {
	e.Class = append(e.Class, classes...)
	return e
}

// WithProperties will merge new attributes to this entity.
// Calling multiple times will continuously append.
func (e *Entity) WithProperties(props Properties) *Entity {
	if e.Properties == nil {
		e.Properties = make(Properties)
	}
	e.Properties.Merge(props)
	return e
}

// WithProperty is a helper for setting a single attribute.
// Calling multiple times will continuously append.
func (e *Entity) WithProperty(key string, value interface{}) *Entity {
	return e.WithProperties(Properties{key: value})
}

// WithLink is a helper for adding a related resource link.
// Calling multiple times will continuously append.
func (e *Entity) WithLink(link *Link) *Entity {
	e.Links = append(e.Links, *link)
	return e
}

// WithAction is a helper for adding a related resource action.
// Calling multiple times will continuously append.
func (e *Entity) WithAction(action *Action) *Entity {
	e.Actions = append(e.Actions, *action)
	return e
}

// Embed is a helper for adding an embedded resource/link to this entity.
// Calling multiple times will continuously append.
func (e *Entity) Embed(entity *EmbeddedEntity) *Entity {
	e.Entities = append(e.Entities, *entity)
	return e
}

// Validate ensures that the entity, embedded entities, links and actions are
// all well-formed.
func (e *Entity) Validate() error {
	for _, ee := range e.Entities {
		if err := validator.Validate(ee); err != nil {
			return err
		}
	}

	for _, l := range e.Links {
		if err := validator.Validate(l); err != nil {
			return err
		}
	}

	for _, a := range e.Actions {
		if err := validator.Validate(a); err != nil {
			return err
		}
	}

	return validator.Validate(e)
}
