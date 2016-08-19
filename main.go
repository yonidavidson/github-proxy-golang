package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["username"])
}

func main() {
	goji.Get("/api/gh/:username", hello)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
