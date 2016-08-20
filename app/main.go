package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"log"
	"net/http"
	"sort"
)

var MODE_DRY bool = false

func getUserData(name string) (Repos, error) {
	var r Repos
	var body []byte
	var err error
	if MODE_DRY {
		body, err = getDataFromLocal()
	} else {
		body, err = getBody(name)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	s := r._map(scorer, nil)
	m := s._map(extractor, []string{"name", "score"})
	sort.Sort(ByScore(m))
	return m, nil
}

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data, err := getUserData(c.URLParams["username"])
	if err != nil {
		fmt.Fprintf(w, "Failed to get data:", 500)
	} else {
		json, _ := json.Marshal(data)
		fmt.Fprintf(w, "%v", string(json))
	}
}

func main() {
	goji.Get("/api/gh/:username", UserHandler)
	goji.Use(middleware.EnvInit)
	goji.Serve()
}
