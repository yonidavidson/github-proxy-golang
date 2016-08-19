package main

import (
	"crypto/tls"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

type Repos []string

func getRepos() ([]byte, error) {
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

func UserHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	repos, err := getRepos()
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
