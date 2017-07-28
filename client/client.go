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
func (c *Client) Get(href string) (*siren.Entity, *http.Response, error) {
	req, err := c.get(href, siren.MediaType)
	if err != nil {
		return nil, nil, err
	}

	return c.request(req)
}

// Follow fetches the entity behind the given siren link.
func (c *Client) Follow(link siren.Link) (*siren.Entity, *http.Response, error) {
	mediaType := link.Type
	if mediaType == "" {
		mediaType = siren.MediaType
	}

	req, err := c.get(string(link.Href), mediaType)
	if err != nil {
		return nil, nil, err
	}

	return c.request(req)
}

// Submit triggers the given action with data supplied by the user.
func (c *Client) Submit(action siren.Action, userData map[string]interface{}) (*siren.Entity, *http.Response, error) {
	u, err := url.Parse(string(action.Href))
	if err != nil {
		return nil, nil, err
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
			body, err = encodeForm(data)
		case "application/json":
			body, err = encodeJSON(data)
		}
		if err != nil {
			return nil, nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("accept", siren.MediaType)
	req.Header.Set("content-type", action.GetType())

	return c.request(req)
}

func (c *Client) request(req *http.Request) (*siren.Entity, *http.Response, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if res.Header.Get("content-type") == siren.MediaType {
		entity, err := c.decodeEntity(res)
		if err != nil {
			return nil, res, err
		}
		return entity, res, nil
	} else {
		return nil, res, nil
	}
}

func (c *Client) get(url, mediaType string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if mediaType != "" {
		req.Header.Set("accept", mediaType)
	}

	return req, nil
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

func (c *Client) decodeEntity(res *http.Response) (*siren.Entity, error) {
	defer res.Body.Close()

	var entity siren.Entity
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&entity); err != nil {
		return nil, ErrInvalidSirenEntity
	}
	return &entity, nil
}

func encodeForm(data map[string]interface{}) (io.Reader, error) {
	q := url.Values{}
	for key, value := range data {
		q.Set(key, fmt.Sprintf("%v", value))
	}
	return strings.NewReader(q.Encode()), nil
}

func encodeJSON(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
