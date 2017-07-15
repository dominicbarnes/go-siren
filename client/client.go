package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	siren "github.com/dominicbarnes/go-siren"
)

// Client is used for interacting with a siren HTTP API.
type Client struct {
	http *http.Client
}

// New creates a new siren client.
func New() *Client {
	return &Client{http: new(http.Client)}
}

// Get retrieves the entity at the given href. This is generally used for the
// entry-point of your application, so prefer using Follow subsequently as your
// user navigates the API.
func (c *Client) Get(href string) (*siren.Entity, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}

	return c.entity(req)
}

// Follow fetches the entity behind the given siren link.
func (c *Client) Follow(link siren.Link) (*siren.Entity, error) {
	return c.Get(string(link.Href))
}

// Submit triggers the given action with data supplied by the user.
func (c *Client) Submit(action siren.Action, userData map[string]interface{}) (*siren.Entity, error) {
	u, err := url.Parse(string(action.Href))
	if err != nil {
		return nil, err
	}

	var body io.Reader
	method := action.GetMethod()
	data := c.data(action, userData)
	if method == http.MethodGet {
		q := u.Query()
		for key, value := range data {
			q.Set(key, fmt.Sprintf("%v", value))
		}
		u.RawQuery = q.Encode()
	} else {
		switch action.GetType() {
		case "application/x-www-form-urlencoded":
			q := url.Values{}
			for key, value := range data {
				q.Set(key, fmt.Sprintf("%v", value))
			}
			body = strings.NewReader(q.Encode())
		case "application/json":
			b, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = bytes.NewBuffer(b)
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", action.GetType())

	return c.entity(req)
}

func (c *Client) data(action siren.Action, userData map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	for _, field := range action.Fields {
		data[field.Name] = field.Value
	}

	for key, value := range userData {
		data[key] = value
	}

	return data
}

func (c *Client) entity(req *http.Request) (*siren.Entity, error) {
	req.Header.Set("accept", siren.MediaType)

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
