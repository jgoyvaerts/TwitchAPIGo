package twitch

import (
	"encoding/json"
	"io/ioutil"
)

type Ingest struct {
	Name         string  `json:"name"`
	Default      bool    `json:"default"`
	ID           int64   `json:"_id"`
	URLTemplate  string  `json:"url_template"`
	Availability float64 `json:"availability"`
}

func (t Twitch) GetIngests() ([]Ingest, error) {
	var ingests []Ingest

	resp, err := t.SendRequest("GET", "/ingests", "v3", Options{})
	if err != nil {
		return ingests, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ingests, err
	}

	type all struct {
		Ingests []Ingest `json:"ingests"`
	}
	var allIngests all

	err = json.Unmarshal(body, &allIngests)
	if err != nil {
		return ingests, err
	}

	return allIngests.Ingests, err
}
