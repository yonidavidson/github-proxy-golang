package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

type Repos []string

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `["yonidavidson.github.io", "DUMMY2"]`)
}

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
