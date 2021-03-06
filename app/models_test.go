package main

import (
	"encoding/json"
	"sort"
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

func TestRepoScoring(t *testing.T) {
	predicat := func(forks float64, stargazers float64, watchers float64) float64 {
		return (forks + 2*stargazers + watchers)
	}
	case1 := Repo{
		"forks_count":      float64(1),
		"stargazers_count": float64(1),
		"watchers_count":   float64(1),
	}
	result := case1.score()
	expected := predicat(case1["forks_count"].(float64), case1["stargazers_count"].(float64), case1["watchers_count"].(float64))
	if result["score"] != expected {
		t.Errorf("scoring failed, expcted: %v , got: %v", expected, result["score"])
	}
}

func TestReposScoring(t *testing.T) {
	var r Repos
	f, _ := getDataFromLocal()
	json.Unmarshal(f, &r)
	m := r._map(scorer, nil)
	for _, o := range m {
		expected := o.score()["score"]
		score, _ := o["score"]
		if expected != score {
			t.Errorf("scoring failed, expcted: %v , got: %v", expected, score)
		}
	}
}

func TestReposSorting(t *testing.T) {
	r := Repos{
		Repo{"score": float64(5)},
		Repo{"score": float64(7)},
		Repo{"score": float64(1)},
		Repo{"score": float64(3)},
	}
	ro := Repos{
		Repo{"score": float64(1)},
		Repo{"score": float64(3)},
		Repo{"score": float64(5)},
		Repo{"score": float64(7)},
	}
	sort.Sort(ByScore(r))
	for i, v := range r {
		if ro[i]["score"] != v["score"] {
			t.Errorf("sorting failed,index:%v ,expcted: %v , got: %v", i, ro[i], r[i])
		}
	}

}
