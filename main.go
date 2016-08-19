package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	repos, err := GetRepos(c.URLParams["username"])
	if err != nil {
		fmt.Fprintf(w, "Failed to get data:", 500)
	} else {
		fmt.Fprintf(w, "%v", repos)
	}
}

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
