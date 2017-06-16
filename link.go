package siren

import validator "gopkg.in/validator.v2"

// Link is a pointer to a related resource.
type Link struct {
	Rel   Rels    `json:"rel" validate:"nonzero"`
	Href  Href    `json:"href" validate:"nonzero"`
	Type  string  `json:"type,omitempty"`
	Title string  `json:"title,omitempty"`
	Class Classes `json:"class,omitempty"`
}

// Validate ensures that the link is well-formed.
func (l Link) Validate() error {
	return validator.Validate(l)
}

// WithBaseHref returns a copy of this link that applies the supplied base
// href to the href and rels.
func (l Link) WithBaseHref(base Href) Link {
	return Link{
		Rel:   l.Rel.WithBaseHref(base),
		Href:  l.Href.WithBaseHref(base),
		Type:  l.Type,
		Title: l.Title,
		Class: l.Class,
	}
}
