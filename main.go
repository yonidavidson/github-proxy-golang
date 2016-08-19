package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	repos, err := GetRepos()
	if err != nil {
		fmt.Fprintf(w, string(repos))
	}
	fmt.Fprintf(w, "Failed to get data", 500)
}

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
