package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

type Repos []string

func GetRepos() ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: nil},
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get("https://api.github.com/users/yonidavidson/repos")
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	return body, err
}
