package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Repos []map[string]interface{}

func (a Repos) _map(tokens ...string) Repos {
	b := a[:0]
	for _, x := range a {
		p := make(map[string]interface{})
		for _, t := range tokens {
			p[t] = x[t]
		}
		b = append(b, p)
	}
	return b
}

func getDataFromLocal() ([]byte, error) {
	file, e := ioutil.ReadFile("./body_test.json")
	return file, e
}

func getBody(name string) ([]byte, error) {
	if MODE_DRY {
		return getDataFromLocal()
	}
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

func getRepos(name string) (Repos, error) {
	var r Repos
	body, err := getBody(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r._map("name"), nil
}
