package siren

import (
	validator "gopkg.in/validator.v2"
)

// EmbeddedEntity is a resource/link that is embedded within a parent entity.
// An embedded resource may contain all the same attributes as any other Entity,
// but also include a Rel indicating it's relationship to the parent. An
// embedded link will only contain attributes that other links have. (eg: href,
// rel, type, class, title)
type EmbeddedEntity struct {
	Entity
	Rel  Rels `json:"rel" validate:"nonzero"`
	Href Href `json:"href,omitempty"`
}

// Validate ensures that the embedded entity is well-formed.
func (e EmbeddedEntity) Validate() error {
	return validator.Validate(e)
}

// WithBaseHref returns a copy of this link that applies the supplied base
// href to the href and rels.
func (e EmbeddedEntity) WithBaseHref(base Href) EmbeddedEntity {
	return EmbeddedEntity{
		Entity: e.Entity.WithBaseHref(base),
		Rel:    e.Rel.WithBaseHref(base),
		Href:   e.Href.WithBaseHref(base),
	}
}
