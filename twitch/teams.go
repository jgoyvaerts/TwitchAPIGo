package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Team struct {
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Background  string    `json:"background"`
	Banner      string    `json:"banner"`
	Logo        string    `json:"logo"`
	ID          int64     `json:"_id"`
	Info        string    `json:"info"`
	DisplayName string    `json:"display_name"`
}

func (t Twitch) GetTeams(page int32) ([]Team, error) {
	var teams []Team

	resp, err := t.SendRequest("GET", "/teams", "v3", t.BuildPageArgs(page))
	if err != nil {
		return teams, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return teams, err
	}

	type allTeams struct {
		Teams []Team `json:"teams"`
	}
	var all allTeams

	err = json.Unmarshal(body, &all)
	return all.Teams, err
}

func (t Twitch) GetTeam(teamName string) (Team, error) {
	var team Team

	resp, err := t.SendRequest("GET", "/teams/"+teamName, "v3", Options{})
	if err != nil {
		return team, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return team, err
	}

	err = json.Unmarshal(body, &team)
	return team, err
}
