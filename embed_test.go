package siren

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmbeddedEntityValidate(t *testing.T) {
	e := EmbeddedEntity{
		Rel:  Rels{"item"},
		Href: "/users/1",
	}
	require.NoError(t, e.Validate())
}

func TestEmbeddedEntityValidateRel(t *testing.T) {
	e := EmbeddedEntity{
		// Rel:  Rels{"item"},
		Href: "/users/1",
	}
	require.Error(t, e.Validate())
}

func TestEmbeddedEntityWithBaseHref(t *testing.T) {
	e := EmbeddedEntity{
		Rel:  Rels{"self", "/rels/custom"},
		Href: "/",
	}
	expected := EmbeddedEntity{
		Rel:  Rels{"self", "https://api.example.com/rels/custom"},
		Href: "https://api.example.com/",
	}
	actual := e.WithBaseHref("https://api.example.com")
	require.EqualValues(t, expected, actual)
}
