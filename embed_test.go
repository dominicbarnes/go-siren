package siren_test

import (
	"errors"
	"testing"

	. "github.com/dominicbarnes/go-siren"

	"github.com/stretchr/testify/require"
)

func TestEmbeddedEntityValidate(t *testing.T) {
	RunValidatorTests(t, map[string]ValidatorSpec{
		"valid": {
			Input: EmbeddedEntity{
				Rel:  Rels{"item"},
				Href: "/users/1",
			},
		},
		"missing rel": {
			Input:    EmbeddedEntity{Href: "/users/1"},
			Expected: errors.New("Rel: zero value"),
		},
	})
}

func TestEmbeddedEntityWithBaseHref(t *testing.T) {
	type spec struct {
		input    EmbeddedEntity
		base     string
		expected EmbeddedEntity
	}

	specs := map[string]spec{
		"rels and href that begin with /": spec{
			input: EmbeddedEntity{
				Rel:  Rels{"/rels/custom"},
				Href: Href("/"),
			},
			base: "https://api.example.com",
			expected: EmbeddedEntity{
				Rel:  Rels{"https://api.example.com/rels/custom"},
				Href: Href("https://api.example.com/"),
			},
		},
		"rels and href that are absolute": spec{
			input: EmbeddedEntity{
				Rel:  Rels{"https://api.example.com/rels/custom"},
				Href: Href("https://api.example.com/"),
			},
			base: "https://example.com",
			expected: EmbeddedEntity{
				Rel:  Rels{"https://api.example.com/rels/custom"},
				Href: Href("https://api.example.com/"),
			},
		},
		"rels that are plain": spec{
			input: EmbeddedEntity{
				Rel:  Rels{"index"},
				Href: Href("https://api.example.com/"),
			},
			base: "https://example.com",
			expected: EmbeddedEntity{
				Rel:  Rels{"index"},
				Href: Href("https://api.example.com/"),
			},
		},
	}

	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			actual := spec.input.WithBaseHref(Href(spec.base))
			require.EqualValues(t, spec.expected, actual)
		})
	}
}
