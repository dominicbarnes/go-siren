package client

import (
	"encoding/json"
	"net/http"

	siren "github.com/dominicbarnes/go-siren"
)

type Client struct {
	http *http.Client
}

func New(client *http.Client) *Client {
	return &Client{http: client}
}

func (c *Client) Get(href string) (*siren.Entity, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}

	return c.entity(req)
}

func (c *Client) Follow(link siren.Link) (*siren.Entity, error) {
	return c.Get(string(link.Href))
}

func (c *Client) Submit(action siren.Action, data map[string]interface{}) (*siren.Entity, error) {
	// TODO: encode request based on type

	req, err := http.NewRequest(action.Method, string(action.Href), nil)
	if err != nil {
		return nil, err
	}

	return c.entity(req)
}

func (c *Client) entity(req *http.Request) (*siren.Entity, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	} else if res.Header.Get("content-type") != siren.MediaType {
		return nil, ErrInvalidMediaType
	}

	var entity siren.Entity
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&entity); err != nil {
		return nil, ErrInvalidSirenEntity
	}
	return &entity, nil
}
