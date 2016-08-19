package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

type envData struct {
	c *web.C
	h http.Handler
}

func (e envData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e.c.Env == nil {
		e.c.Env = make(map[interface{}]interface{})
	}
	e.c.Env["data"] = "githubData"
	e.h.ServeHTTP(w, r)
}

func GetData(c *web.C, h http.Handler) http.Handler {
	return envData{c, h}
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["username"])
}

func main() {
	goji.Get("/api/gh/:username", hello)
	goji.Use(GetData)
	goji.Serve()
}
