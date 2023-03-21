package cyoa

import (
	"encoding/json"
	"io/ioutil"
)

type Arc struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []ArcOption `json:"options"`
}

type ArcOption struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}

// GetArcMap returns a map of arcs from the sample arc JSON file, "gopher.json"
func GetArcMap() (map[string]Arc, error) {
	jsn, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		return nil, err
	}

	arcMap := make(map[string]Arc)
	err = json.Unmarshal(jsn, &arcMap)
	if err != nil {
		return nil, err
	}

	return arcMap, err
}
