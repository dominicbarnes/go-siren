package siren

import "gopkg.in/validator.v2"

// EmbeddedEntity is a resource/link that is embedded within a parent entity.
// An embedded resource may contain all the same attributes as any other Entity,
// but also include a Rel indicating it's relationship to the parent. An
// embedded link will only contain attributes that other links have. (eg: href,
// rel, type, class, title)
type EmbeddedEntity struct {
	Rel  []string `json:"rel" validate:"nonzero"`
	Href string   `json:"href,omitempty"`
	Entity
}

// NewEmbeddedEntity is a helper for creating a new EmbeddedEntity instance that
// represents an embedded resource.
func NewEmbeddedEntity(rel []string) *EmbeddedEntity {
	e := NewEntity("")
	return &EmbeddedEntity{
		Entity: *e,
		Rel:    rel,
	}
}

// NewEmbeddedLink is a helper for creating a new EmbeddedEntity instance that
// represents an embedded link.
func NewEmbeddedLink(rel []string, href string) *EmbeddedEntity {
	e := NewEntity("")
	return &EmbeddedEntity{
		Entity: *e,
		Rel:    rel,
		Href:   href,
	}
}

// WithTitle is a ported version of Entity.WithTitle
func (e *EmbeddedEntity) WithTitle(title string) *EmbeddedEntity {
	e.Entity.WithTitle(title)
	return e
}

// WithClasses is a ported version of Entity.WithClasses
func (e *EmbeddedEntity) WithClasses(classes ...string) *EmbeddedEntity {
	e.Entity.WithClasses(classes...)
	return e
}

// WithProperties is a ported version of Entity.WithProperties
func (e *EmbeddedEntity) WithProperties(props Properties) *EmbeddedEntity {
	e.Entity.WithProperties(props)
	return e
}

// WithProperty is a ported version of Entity.WithProperty
func (e *EmbeddedEntity) WithProperty(key string, value interface{}) *EmbeddedEntity {
	e.Entity.WithProperty(key, value)
	return e
}

// WithLink is a ported version of Entity.WithLink
func (e *EmbeddedEntity) WithLink(link *Link) *EmbeddedEntity {
	e.Entity.WithLink(link)
	return e
}

// WithAction is a ported version of Entity.WithAction
func (e *EmbeddedEntity) WithAction(action *Action) *EmbeddedEntity {
	e.Entity.WithAction(action)
	return e
}

// Validate ensures the embedded entity/link is well-formed.
func (e *EmbeddedEntity) Validate() error {
	return validator.Validate(e)
}
