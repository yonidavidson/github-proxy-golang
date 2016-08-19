package main

import (
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
	MODE_DRY = true
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
}
