package main

import (
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"sort"
)

var REPO_ATTRIBUTES = []string{"name", "score", "created_at", "language",
	"forks", "watchers", "open_issues_count", "stargazers_count"}

func getUserData(name string) (Repos, error) {
	var r Repos
	var body []byte
	var err error
	if MODE_DRY == "TRUE" {
		log.Println("warning - dry mode, data from localhost")
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
	m := s._map(extractor, REPO_ATTRIBUTES)
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
