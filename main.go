package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

var MODE_DRY bool = false

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	repos, err := getRepos(c.URLParams["username"])
	if err != nil {
		fmt.Fprintf(w, "Failed to get data:", 500)
	} else {
		json, _ := json.Marshal(repos)
		fmt.Fprintf(w, "%v", string(json))
	}
}

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
