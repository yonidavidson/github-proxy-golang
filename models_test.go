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

	m := r._map(extractor,
		[]string{"name", "stargazers_count", "forks",
			"watchers", "forks_count", "stargazers_count", "watchers_count"})
	for _, o := range m {
		_, ok1 := o["name"]
		_, ok2 := o["stargazers_count"]
		_, ok3 := o["forks"]
		_, ok4 := o["watchers"]
		_, ok5 := o["forks_count"]
		_, ok6 := o["stargazers_count"]
		_, ok7 := o["watchers_count"]
		if !(ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7) {
			t.Errorf("mapping error: %v\n", o)
		}
		t.Log("%v", m)
	}
}

func TestMappingNegative(t *testing.T) {
	var r Repos
	f, _ := getDataFromLocal()
	json.Unmarshal(f, &r)
	m := r._map(extractor, []string{"name", "THIS_IS_WRONG"})
	for _, o := range m {
		_, ok1 := o["name"]
		_, ok2 := o["THIS_IS_WRONG"]
		if ok1 && ok2 {
			t.Errorf("mapping error - should have failed: %v\n", o)
		}
		t.Log("%v", m)
	}
}

func TestScoring(t *testing.T) {
	predicat := func(forks int, stargazers int, watchers int) int { return (forks + 2*stargazers + watchers) }
	case1 := Repo{
		"forks_count":      1,
		"stargazers_count": 1,
		"watchers_count":   1,
	}
	result := case1.score()
	expected := predicat(case1["forks_count"].(int), case1["stargazers_count"].(int), case1["watchers_count"].(int))
	if result["score"] != expected {
		t.Errorf("scoring failed, expcted: %v , got: %v", expected, result["score"])
	}
}
