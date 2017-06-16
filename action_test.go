package siren

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestActionValidate(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "/search",
	}
	require.NoError(t, a.Validate())
}

func TestActionValidateName(t *testing.T) {
	a := Action{
		// Name: "search",
		Href: "/search",
	}
	require.Error(t, a.Validate())
}

func TestActionValidateHref(t *testing.T) {
	a := Action{
		Name: "search",
		// Href: "/search",
	}
	require.Error(t, a.Validate())
}

func TestActionWithBaseHref(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "/search",
	}
	expected := Action{
		Name: "search",
		Href: "https://api.example.com/search",
	}
	actual := a.WithBaseHref("https://api.example.com")
	require.EqualValues(t, expected, actual)
}

func TestActionWithBaseHrefAbsolute(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "https://api.example.com/search",
	}
	expected := Action{
		Name: "search",
		Href: "https://api.example.com/search",
	}
	actual := a.WithBaseHref("https://example.com")
	require.EqualValues(t, expected, actual)
}
