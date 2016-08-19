package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func getData(t *testing.T) []byte {
	file, e := ioutil.ReadFile("./body_test.json")
	if e != nil {
		t.Errorf("File error: %v\n", e)
	}
	return file
}

func getDataMap(t *testing.T) Repos {
	var r Repos
	f := getData(t)
	e := json.Unmarshal(f, &r)
	if e != nil {
		t.Errorf("Unmarshling error: %v\n", e)
	}
	return r
}

func TestUnmarsh(t *testing.T) {
	getDataMap(t)
}

func TestMapping(t *testing.T) {
	m := getDataMap(t)
	f := m._map("name", "stargazers_count")
	fmt.Printf("%v", f)
}

// func TestNestedMapping(t *testing.T) {
// 	m := getDataMap(t)
// 	f = m.map("name", "stargazers_count")
// }

// func TestRepoMap(t *testing.T) {
// 	f := getData(t)

// }
