package siren_test

import (
	"errors"
	"net/http"
	"testing"

	. "github.com/dominicbarnes/go-siren"

	"github.com/stretchr/testify/require"
)

func TestActionValidate(t *testing.T) {
	RunValidatorTests(t, map[string]ValidatorSpec{
		"valid": {
			Input: Action{Name: "search", Href: "/search"},
		},
		"missing name": {
			Input:    Action{Href: "/search"},
			Expected: errors.New("Name: zero value"),
		},
		"missing href": {
			Input:    Action{Name: "search"},
			Expected: errors.New("Href: zero value"),
		},
	})
}

func TestActionWithBaseHref(t *testing.T) {
	type Spec struct {
		action   Action
		base     string
		expected Action
	}

	specs := map[string]Spec{
		"resolve relative href": {
			action: Action{
				Name: "search",
				Href: "/search",
			},
			base: "https://api.example.com",
			expected: Action{
				Name: "search",
				Href: "https://api.example.com/search",
			},
		},
		"ignore absolute href": {
			action: Action{
				Name: "search",
				Href: "https://api.example.com/search",
			},
			base: "https://example.com",
			expected: Action{
				Name: "search",
				Href: "https://api.example.com/search",
			},
		},
	}

	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			actual := spec.action.WithBaseHref(Href(spec.base))
			require.EqualValues(t, spec.expected, actual)
		})
	}
}

func TestActionGetMethod(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		a := Action{}
		require.Equal(t, ActionDefaultMethod, a.GetMethod())
	})

	t.Run("specified", func(t *testing.T) {
		a := Action{Method: http.MethodPost}
		require.Equal(t, a.Method, a.GetMethod())
	})
}

func TestActionGetType(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		a := Action{}
		require.Equal(t, ActionDefaultType, a.GetType())
	})

	t.Run("specified", func(t *testing.T) {
		a := Action{Type: "application/json"}
		require.Equal(t, a.Type, a.GetType())
	})
}
