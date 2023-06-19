package siren_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	. "github.com/dominicbarnes/go-siren"

	"github.com/stretchr/testify/require"
)

func TestEntityValidate(t *testing.T) {
	RunValidatorTests(t, map[string]ValidatorSpec{
		"empty": {
			Input: Entity{},
		},
		"embed valid": {
			Input: Entity{
				Entities: []EmbeddedEntity{
					{Href: "/posts/1", Rel: Rels{"item"}},
				},
			},
		},
		"embed missing rel": {
			Input: Entity{
				Entities: []EmbeddedEntity{
					{},
				},
			},
			Expected: errors.New("Rel: zero value"),
		},
		"link valid": {
			Input: Entity{
				Links: []Link{
					{Href: "/", Rel: Rels{"self"}},
				},
			},
		},
		"link missing rel": {
			Input: Entity{
				Links: []Link{
					{Href: "/"},
				},
			},
			Expected: errors.New("Rel: zero value"),
		},
		"link missing href": {
			Input: Entity{
				Links: []Link{
					{Rel: Rels{"self"}},
				},
			},
			Expected: errors.New("Href: zero value"),
		},
		"action valid": {
			Input: Entity{
				Actions: []Action{
					{Name: "search", Href: "/search"},
				},
			},
		},
		"action missing name": {
			Input: Entity{
				Actions: []Action{
					{Href: "/search"},
				},
			},
			Expected: errors.New("Name: zero value"),
		},
		"action missing href": {
			Input: Entity{
				Actions: []Action{
					{Name: "search"},
				},
			},
			Expected: errors.New("Href: zero value"),
		},
	})

	t.Run("WithBaseHref()", func(t *testing.T) {
		e := Entity{
			Entities: []EmbeddedEntity{
				{Href: "/posts/1", Rel: Rels{"item"}},
			},
			Links: []Link{
				{Href: "/", Rel: Rels{"self"}},
			},
			Actions: []Action{
				{Name: "search", Href: "/search"},
			},
		}
		expected := Entity{
			Entities: []EmbeddedEntity{
				{Href: "https://api.example.com/posts/1", Rel: Rels{"item"}},
			},
			Links: []Link{
				{Href: "https://api.example.com/", Rel: Rels{"self"}},
			},
			Actions: []Action{
				{Name: "search", Href: "https://api.example.com/search"},
			},
		}
		actual := e.WithBaseHref("https://api.example.com")
		require.EqualValues(t, expected, actual)
	})
}

func ExampleEntity() {
	e := Entity{
		Class: Classes{"order"},
		Properties: Properties{
			"orderNumber": 42,
			"itemCount":   3,
			"status":      "pending",
		},
		Entities: []EmbeddedEntity{
			{
				Rel:  Rels{"http://x.io/rels/order-items"},
				Href: "/orders/42/items",
				Entity: Entity{
					Class: Classes{"items", "collection"},
				},
			},
			{
				Rel:  Rels{"http://x.io/rels/customer"},
				Href: "/orders/42/items",
				Entity: Entity{
					Class: Classes{"info", "customer"},
					Properties: Properties{
						"customerId": "pj123",
						"name":       "Peter Joseph",
					},
					Links: []Link{
						{Rel: Rels{"self"}, Href: "/customers/pj123"},
					},
				},
			},
		},
		Actions: []Action{
			{
				Name:   "add-item",
				Title:  "Add Item",
				Method: "POST",
				Href:   "/orders/42/items",
				Type:   "application/x-www-form-urlencoded",
				Fields: []ActionField{
					{Name: "orderNumber", Type: "hidden", Value: 42},
					{Name: "productCode", Type: "text"},
					{Name: "quantity", Type: "number"},
				},
			},
		},
		Links: []Link{
			{Rel: Rels{"self"}, Href: "/orders/42"},
			{Rel: Rels{"previous"}, Href: "/orders/41"},
			{Rel: Rels{"next"}, Href: "/orders/43"},
		},
	}

	data, err := json.MarshalIndent(e.WithBaseHref("http://api.x.io"), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
