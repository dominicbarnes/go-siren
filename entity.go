package siren

import (
	"gopkg.in/validator.v2"
)

// Entity is a top-level resource in a siren API.
type Entity struct {
	Entities   []EmbeddedEntity `json:"entities,omitempty"`
	Links      []Link           `json:"links,omitempty"`
	Actions    []Action         `json:"actions,omitempty"`
	Properties Properties       `json:"properties,omitempty"`
	Title      string           `json:"title,omitempty"`
	Class      Classes          `json:"class,omitempty"`
}

// Validate ensures that the entity, embedded entities, links and actions are
// all well-formed.
func (e Entity) Validate() error {
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

// WithBaseHref applies the given base href to the sub-entities, links and
// actions.
func (e Entity) WithBaseHref(base Href) Entity {
	var entities []EmbeddedEntity
	for _, embed := range e.Entities {
		entities = append(entities, embed.WithBaseHref(base))
	}

	var links []Link
	for _, link := range e.Links {
		links = append(links, link.WithBaseHref(base))
	}

	var actions []Action
	for _, action := range e.Actions {
		actions = append(actions, action.WithBaseHref(base))
	}

	return Entity{
		Entities:   entities,
		Links:      links,
		Actions:    actions,
		Properties: e.Properties,
		Title:      e.Title,
		Class:      e.Class,
	}
}
