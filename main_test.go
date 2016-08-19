package main

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestGetHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/gh/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := web.HandlerFunc(UserHandler)

	urlparams := make(map[string]string)
	urlparams["username"] = "yoni"
	ctx := web.C{urlparams, make(map[interface{}]interface{})}

	handler.ServeHTTPC(ctx, rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var data Repos
	expected := "[item1, item2 ....]"
	json.Unmarshal([]byte(rr.Body.String()), &data)
	if len(data) < 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	expected = "yonidavidson.github.io"
	if !contains(data, expected) {
		t.Errorf("handler returned unexpected body: got %v did not contain '[%v]'",
			rr.Body.String(), expected)
	}
}
