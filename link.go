package siren

import "gopkg.in/validator.v2"

// Link is a pointer to a related resource.
type Link struct {
	Rel   Rels    `json:"rel" validate:"nonzero"`
	Href  string  `json:"href" validate:"nonzero"`
	Type  string  `json:"type,omitempty"`
	Title string  `json:"title,omitempty"`
	Class Classes `json:"class,omitempty"`
}

// NewLink is a helper for creating a new Link instance.
func NewLink(rel Rels, href string) *Link {
	return &Link{Rel: rel, Href: href}
}

// WithType is a helper for setting an optional media type for this linked resource.
// Calling multiple times will replace any previous values.
func (l *Link) WithType(mediaType string) *Link {
	l.Type = mediaType
	return l
}

// WithTitle is a helper for setting an optional title for this linked resource.
// Calling multiple times will replace any previous values.
func (l *Link) WithTitle(title string) *Link {
	l.Title = title
	return l
}

// WithClasses appends new class names to the link.
// Calling multiple times will continuously append.
func (l *Link) WithClasses(classes ...string) *Link {
	l.Class = append(l.Class, classes...)
	return l
}

// Validate ensures the link is well-formed.
func (l *Link) Validate() error {
	return validator.Validate(l)
}
