package siren_test

import (
	"errors"
	"testing"

	. "github.com/dominicbarnes/go-siren"

	"github.com/stretchr/testify/require"
)

func TestLinkValidate(t *testing.T) {
	RunValidatorTests(t, map[string]ValidatorSpec{
		"valid": {
			Input: Link{Href: "/", Rel: Rels{"self"}},
		},
		"missing href": {
			Input:    Link{Rel: Rels{"self"}},
			Expected: errors.New("Href: zero value"),
		},
		"missing rel": {
			Input:    Link{Href: "/"},
			Expected: errors.New("Rel: zero value"),
		},
	})
}

func TestLinkWithBaseHref(t *testing.T) {
	l := Link{
		Href: "/",
		Rel:  Rels{"self", "/rels/custom"},
	}
	expected := Link{
		Href: "https://api.example.com/",
		Rel:  Rels{"self", "https://api.example.com/rels/custom"},
	}
	actual := l.WithBaseHref("https://api.example.com")
	require.EqualValues(t, expected, actual)
}
