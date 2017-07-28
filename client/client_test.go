package client_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	siren "github.com/dominicbarnes/go-siren"
	. "github.com/dominicbarnes/go-siren/client"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client *Client
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupSuite() {
	suite.client = New()
}

func (suite *ClientTestSuite) TestGet() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodGet, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal(siren.MediaType, r.Header.Get("accept"))

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Get(ts.URL + "/entity")
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestGetInvalidMediaType() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// send an invalid response for the client
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte("hello world"))
	}))

	entity, res, err := suite.client.Get(ts.URL)
	suite.NoError(err)
	suite.NotNil(res)
	suite.Nil(entity)

	body, err := ioutil.ReadAll(res.Body)
	suite.NoError(err)
	suite.EqualValues("hello world", body)
}

func (suite *ClientTestSuite) TestGetInvalidURL() {
	entity, res, err := suite.client.Get(":")
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestGetEmptyURL() {
	entity, res, err := suite.client.Get("")
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestGetInvalidSirenEntity() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// send an invalid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{"class":1}`)) // will not unmarshal
	}))

	entity, res, err := suite.client.Get(ts.URL)
	suite.EqualValues(err, ErrInvalidSirenEntity)
	suite.NotNil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestFollow() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodGet, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal(siren.MediaType, r.Header.Get("accept"))

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Follow(siren.Link{
		Href: siren.Href(ts.URL + "/entity"),
		Rel:  siren.Rels{"self"},
	})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestFollowEmptyHref() {
	entity, res, err := suite.client.Follow(siren.Link{
		Href: siren.Href(""),
		Rel:  siren.Rels{"invalid"},
	})
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestFollowInvalidHref() {
	entity, res, err := suite.client.Follow(siren.Link{
		Href: siren.Href(":"),
		Rel:  siren.Rels{"invalid"},
	})
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestSubmit() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPost, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal(siren.MediaType, r.Header.Get("accept"))

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(ts.URL + "/entity"),
	}, nil)
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitGetQuery() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodGet, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("foo=bar", r.URL.Query().Encode())
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name: "do-stuff",
		Href: siren.Href(ts.URL + "/entity"),
	}, map[string]interface{}{"foo": "bar"})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitPostForm() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPost, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.EqualValues("foo=bar", body)

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(ts.URL + "/entity"),
	}, map[string]interface{}{"foo": "bar"})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitPostFormExplicit() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPost, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.EqualValues("foo=bar", body)

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(ts.URL + "/entity"),
		Type:   "application/x-www-form-urlencoded",
	}, map[string]interface{}{"foo": "bar"})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitPatchJSON() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPatch, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/json", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.JSONEq(`{"foo":"bar"}`, string(body))

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPatch,
		Href:   siren.Href(ts.URL + "/entity"),
		Type:   "application/json",
	}, map[string]interface{}{"foo": "bar"})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitNoData() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodDelete, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.Empty(body)

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodDelete,
		Href:   siren.Href(ts.URL + "/entity"),
	}, nil)
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitDefaultData() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPost, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.EqualValues("foo=bar", body)

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(ts.URL + "/entity"),
		Fields: []siren.ActionField{
			{Name: "foo", Value: "bar"},
		},
	}, nil)
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitDefaultDataWithUserData() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assert expected request was sent
		suite.Equal(http.MethodPost, r.Method)
		suite.Equal("/entity", r.URL.Path)
		suite.Equal("application/x-www-form-urlencoded", r.Header.Get("content-type"))
		suite.Equal(siren.MediaType, r.Header.Get("accept"))
		body, err := ioutil.ReadAll(r.Body)
		suite.NoError(err)
		suite.EqualValues("foo=baz", body)

		// send a valid response for the client
		w.Header().Set("content-type", siren.MediaType)
		w.Write([]byte(`{}`))
	}))

	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(ts.URL + "/entity"),
		Fields: []siren.ActionField{
			{Name: "foo", Value: "bar"},
		},
	}, map[string]interface{}{"foo": "baz"})
	suite.NoError(err)
	suite.NotNil(res)
	suite.EqualValues(entity, new(siren.Entity))
}

func (suite *ClientTestSuite) TestSubmitInvalidHref() {
	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(":"),
		Fields: []siren.ActionField{
			{Name: "foo", Value: "bar"},
		},
	}, nil)
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}

func (suite *ClientTestSuite) TestSubmitEmptyHref() {
	entity, res, err := suite.client.Submit(siren.Action{
		Name:   "do-stuff",
		Method: http.MethodPost,
		Href:   siren.Href(""),
		Fields: []siren.ActionField{
			{Name: "foo", Value: "bar"},
		},
	}, nil)
	suite.Error(err)
	suite.Nil(res)
	suite.Nil(entity)
}
