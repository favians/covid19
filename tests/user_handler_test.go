package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/favians/golang_starter/router"

	"github.com/gavv/httpexpect"
)

func TestGetUsers(t *testing.T) {

	RebuildData()

	// create http.Handler
	handler := router.New()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	//get normal token
	obj := e.GET("/login").
		WithQuery("username", "vian").WithQuery("password", "1234").
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	// without query string
	obj = auth.GET("/users").WithQuery("id", "1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ValueEqual("id", 1)
	obj.ValueEqual("name", "vian")
	obj.ValueEqual("username", "vian")

	// // with rp & p set
	obj = auth.GET("/users/list").WithQuery("p", "1").WithQuery("rp", "2").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("data").Array().Element(0).Object().Value("id").Equal(1)

	// // with filter name
	obj = auth.GET("/users/list").WithQuery("name", "vian").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().Contains("data")
	obj.Value("data").Array().Element(0).Object().Value("name").Equal("vian")
}

func TestAddUser(t *testing.T) {

	RebuildData()

	// create http.Handler
	handler := router.New()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	obj := e.GET("/login").
		WithQuery("username", "vian").WithQuery("password", "1234").
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	payload := make(map[string]interface{})

	// normal add new
	payload = map[string]interface{}{
		"name":     "unit test user",
		"username": "unit_test_username",
		"password": "unit_test_password",
	}
	obj = auth.POST("/users").
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()
	obj.ContainsKey("name").ValueEqual("name", "unit test user")
	obj.ContainsKey("username").ValueEqual("username", "unit_test_username")

	// failed with 422
	payload = map[string]interface{}{
		"name": "unit test user",
	}
	obj = auth.POST("/users").
		WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()
	obj.Value("validation_errors").Object().Value("password").Array().Element(0).String().Equal("The password field is required")
}
