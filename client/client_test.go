package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	siren "github.com/dominicbarnes/go-siren"
	. "github.com/dominicbarnes/go-siren/client"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupSuite() {
	suite.server = httptest.NewServer(mux())
	suite.client = New(new(http.Client))
}

func (suite *ClientTestSuite) TestGetNotFound() {
	entity, err := suite.client.Get(suite.url("/does-not-exist"))
	suite.EqualValues(err, ErrInvalidMediaType)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestGetNotJSON() {
	entity, err := suite.client.Get(suite.url("/not-json"))
	suite.EqualValues(err, ErrInvalidMediaType)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestGetNotValidSiren() {
	entity, err := suite.client.Get(suite.url("/not-valid-siren"))
	suite.EqualValues(err, ErrInvalidSirenEntity)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestGetSuccess() {
	entity, err := suite.client.Get(suite.url("/sample"))
	suite.NoError(err)
	suite.NotEmpty(entity)
	suite.EqualValues(siren.Classes{"order"}, entity.Class) // not a full equality check, just enough to confirm it was decoded at all
}

func (suite *ClientTestSuite) TestFollow() {
	entity, err := suite.client.Follow(siren.Link{
		Href: siren.Href(suite.url("/sample")),
		Rel:  siren.Rels{"item"},
	})
	suite.NoError(err)
	suite.NotEmpty(entity)
	suite.EqualValues(siren.Classes{"order"}, entity.Class) // not a full equality check, just enough to confirm it was decoded at all
}

func (suite *ClientTestSuite) TestSubmit() {
	action := siren.Action{
		Name: "do-stuff",
		Href: siren.Href(suite.url("/sample")),
	}

	entity, err := suite.client.Submit(action, nil)
	suite.NoError(err)
	suite.NotEmpty(entity)
	suite.EqualValues(siren.Classes{"order"}, entity.Class) // not a full equality check, just enough to confirm it was decoded at all
}

func (suite *ClientTestSuite) url(input string) string {
	return suite.server.URL + input
}

func mux() *http.ServeMux {
	mux := http.NewServeMux()

	entities := map[string]siren.Entity{
		"empty": siren.Entity{},
		"sample": siren.Entity{
			BaseHref: "http://api.x.io",
			Class:    siren.Classes{"order"},
			Properties: siren.Properties{
				"orderNumber": 42,
				"itemCount":   3,
				"status":      "pending",
			},
			Entities: []siren.EmbeddedEntity{
				{
					Rel:  siren.Rels{"http://x.io/rels/order-items"},
					Href: "/orders/42/items",
					Entity: siren.Entity{
						Class: siren.Classes{"items", "collection"},
					},
				},
				{
					Rel:  siren.Rels{"http://x.io/rels/customer"},
					Href: "/orders/42/items",
					Entity: siren.Entity{
						Class: siren.Classes{"info", "customer"},
						Properties: siren.Properties{
							"customerId": "pj123",
							"name":       "Peter Joseph",
						},
						Links: []siren.Link{
							{Rel: siren.Rels{"self"}, Href: "/customers/pj123"},
						},
					},
				},
			},
			Actions: []siren.Action{
				{
					Name:   "add-item",
					Title:  "Add Item",
					Method: "POST",
					Href:   "/orders/42/items",
					Type:   "application/x-www-form-urlencoded",
					Fields: []siren.ActionField{
						{Name: "orderNumber", Type: "hidden", Value: 42},
						{Name: "productCode", Type: "text"},
						{Name: "quantity", Type: "number"},
					},
				},
			},
			Links: []siren.Link{
				{Rel: siren.Rels{"self"}, Href: "/orders/42"},
				{Rel: siren.Rels{"previous"}, Href: "/orders/41"},
				{Rel: siren.Rels{"next"}, Href: "/orders/43"},
			},
		},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.Split(r.URL.Path, "/")[1]
		if entity, ok := entities[key]; ok {
			w.Header().Set("content-type", siren.MediaType)
			json.NewEncoder(w).Encode(entity)
		} else {
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/not-json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte("not json"))
	})

	mux.HandleFunc("/not-siren", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{}"))
	})

	mux.HandleFunc("/not-valid-siren", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{"class":1}`)) // 1 is not a valid class (string array)
	})

	return mux
}
