package siren

// Rels is a collection of link relations. They can either be short names as
// defined by the IANA, or full URLs that are application-specific.
type Rels []Href

// WithBaseHref applies the given base href to all the rels that begin with a /.
// This allows using absolute URLs and IANA predefined rels.
func (r Rels) WithBaseHref(base Href) Rels {
	rels := make(Rels, len(r))
	for x, rel := range r {
		rels[x] = rel.WithBaseHref(base)
	}
	return rels
}
