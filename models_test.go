package main

import (
	"encoding/json"
	"testing"
)

func TestUnmarsh(t *testing.T) {
	var r Repos
	f, _ := getDataFromLocal()
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Errorf("Unmarshling error: %v\n", err)
	}
}

func TestMapping(t *testing.T) {
	var r Repos
	f, _ := getDataFromLocal()
	json.Unmarshal(f, &r)
	m := r._map("name", "stargazers_count")
	for _, o := range m {
		_, ok1 := o["name"]
		_, ok2 := o["stargazers_count"]
		if !(ok1 && ok2) {
			t.Errorf("mapping error: %v\n", o)
		}
	}
}
