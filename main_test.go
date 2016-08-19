package main

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/gh/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := web.HandlerFunc(hello)

	urlparams := make(map[string]string)
	urlparams["username"] = "yoni"
	ctx := web.C{urlparams, make(map[interface{}]interface{})}

	handler.ServeHTTPC(ctx, rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, yoni!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestGetDataHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/gh/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	ctx := web.C{make(map[string]string), make(map[interface{}]interface{})}

	handler := GetData(&ctx, web.HandlerFunc(hello))
	handler.ServeHTTP(rr, req)

	if _, ok := ctx.Env["data"]; !ok {
		t.Errorf("No data was fetched")
	}
}
