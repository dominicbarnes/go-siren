package siren

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntityValidateEmpty(t *testing.T) {
	e := Entity{}
	require.NoError(t, e.Validate())
}

func TestEntityValidateValidEmbed(t *testing.T) {
	e := Entity{
		Entities: []EmbeddedEntity{
			{Href: "/posts/1", Rel: Rels{"item"}},
		},
	}
	require.NoError(t, e.Validate())
}

func TestEntityValidateInvalidEmbed(t *testing.T) {
	e := Entity{
		Entities: []EmbeddedEntity{
			{},
		},
	}
	require.Error(t, e.Validate())
}

func TestEntityValidateValidLink(t *testing.T) {
	e := Entity{
		Links: []Link{
			{Href: "/", Rel: Rels{"self"}},
		},
	}
	require.NoError(t, e.Validate())
}

func TestEntityValidateInvalidLink(t *testing.T) {
	e := Entity{
		Links: []Link{
			{},
		},
	}
	require.Error(t, e.Validate())
}

func TestEntityValidateValidAction(t *testing.T) {
	e := Entity{
		Actions: []Action{
			{Name: "search", Href: "/search"},
		},
	}
	require.NoError(t, e.Validate())
}

func TestEntityValidateInvalidAction(t *testing.T) {
	e := Entity{
		Actions: []Action{
			{},
		},
	}
	require.Error(t, e.Validate())
}

func TestEntityWithBaseHref(t *testing.T) {
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
}

func TestEntityMarshalJSON(t *testing.T) {
	expected, err := json.Marshal(Entity{
		Entities: []EmbeddedEntity{
			{Href: "https://api.example.com/posts/1", Rel: Rels{"item"}},
		},
		Links: []Link{
			{Href: "https://api.example.com/", Rel: Rels{"self"}},
		},
		Actions: []Action{
			{Name: "search", Href: "https://api.example.com/search"},
		},
	})
	require.NoError(t, err)

	actual, err := json.Marshal(Entity{
		BaseHref: "https://api.example.com",
		Entities: []EmbeddedEntity{
			{Href: "/posts/1", Rel: Rels{"item"}},
		},
		Links: []Link{
			{Href: "/", Rel: Rels{"self"}},
		},
		Actions: []Action{
			{Name: "search", Href: "/search"},
		},
	})
	require.NoError(t, err)

	require.EqualValues(t, expected, actual)
}

func ExampleEntity() {
	e := Entity{
		BaseHref: "http://api.x.io",
		Class:    Classes{"order"},
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

	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
