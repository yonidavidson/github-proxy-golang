package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

type Repo map[string]interface{}
type Repos []Repo
type ReposMapper func(Repo, interface{}) Repo

type ByScore Repos

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i]["score"].(float64) < a[j]["score"].(float64) }

var extractor ReposMapper = func(x Repo, data interface{}) Repo {
	tokens := data.([]string)
	p := make(Repo)
	for _, t := range tokens {
		_, ok := x[t]
		if ok {
			p[t] = x[t]
		}
	}
	return p
}

var scorer ReposMapper = func(x Repo, data interface{}) Repo {
	return x.score()
}

func (r Repo) score() Repo {
	//For simplicity i assume that all this fields exist in json (or I would need to check each assigment before)
	forks := r["forks_count"].(float64)
	stargazers := r["stargazers_count"].(float64)
	watchers := r["watchers_count"].(float64)
	r["score"] = forks + 2*stargazers + watchers
	return r
}

func (a Repos) _map(m ReposMapper, d interface{}) Repos {
	b := a[:0]
	for _, x := range a {
		b = append(b, m(x, d))
	}
	return b
}

func getDataFromLocal() ([]byte, error) {
	file, e := ioutil.ReadFile("./body_test.json")
	return file, e
}

func getBody(name string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: nil},
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get("https://api.github.com/users/" + name + "/repos")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
