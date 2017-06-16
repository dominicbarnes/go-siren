package siren

import "strings"

// Href is a wrapper for a string that is a resource href that can automatically
// be prefixed with a base href.
type Href string

// WithBaseHref applies the given base href to this href. This prefix is only
// applied when the href begins with a "/".
func (h Href) WithBaseHref(base Href) Href {
	if strings.HasPrefix(string(h), "/") {
		return base + h
	}

	return h
}
