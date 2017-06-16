package siren

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkValidate(t *testing.T) {
	l := Link{
		Href: "/",
		Rel:  Rels{"self"},
	}
	require.NoError(t, l.Validate())
}

func TestLinkValidateHref(t *testing.T) {
	l := Link{
		// Href: "/",
		Rel: Rels{"self"},
	}
	require.Error(t, l.Validate())
}

func TestLinkValidateRel(t *testing.T) {
	l := Link{
		Href: "/",
		// Rel: Rels{"self"},
	}
	require.Error(t, l.Validate())
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
