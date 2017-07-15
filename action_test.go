package siren

import (
	"net/http"
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

func TestActionGetMethod(t *testing.T) {
	a := Action{
		Name:   "search",
		Href:   "/search",
		Method: http.MethodPost,
	}
	require.Equal(t, http.MethodPost, a.GetMethod())
}

func TestActionGetMethodDefault(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "/search",
	}
	require.Equal(t, ActionDefaultMethod, a.GetMethod())
}

func TestActionGetType(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "/search",
		Type: "application/json",
	}
	require.Equal(t, "application/json", a.GetType())
}

func TestActionGetTypeDefault(t *testing.T) {
	a := Action{
		Name: "search",
		Href: "/search",
	}
	require.Equal(t, ActionDefaultType, a.GetType())
}
