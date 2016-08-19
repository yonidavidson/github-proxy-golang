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
	r._map("name", "stargazers_count")
}
