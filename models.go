package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RepoProps map[string]interface{}
type Repos []RepoProps

func filter(a Repos) Repos {
	// b := a[:0]
	// for _, x := range a {
	// 	// m := x.(map[string]interface{})
	// 	// b = append(b, m["description"])
	// 	x := nil
	// }
	// return b
	return nil
}

func GetRepos() (Repos, error) {
	var r Repos = nil
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: nil},
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get("https://api.github.com/users/yonidavidson/repos")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	e := json.Unmarshal(body, &r)
	if e != nil {
		log.Println(err)
		return nil, err
	}
	n := filter(r)
	return n, err
}
