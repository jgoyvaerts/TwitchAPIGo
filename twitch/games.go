package twitch

import (
	"encoding/json"
	"io/ioutil"
)

type Box struct {
	Large    string `json:"large"`
	Medium   string `json:"medium"`
	Small    string `json:"small"`
	Template string `json:"template"`
}

type Logo struct {
	Large    string `json:"large"`
	Medium   string `json:"medium"`
	Small    string `json:"small"`
	Template string `json:"template"`
}

type Game struct {
	Name        string `json:"name"`
	Box         Box    `json:"box"`
	Logo        Logo   `json:"logo"`
	ID          int64  `json:"_id"`
	GiantBombID int64  `json:"giantbomb_id"`
}

type TopGame struct {
	Game     Game  `json:"game"`
	Viewers  int64 `json:"viewers"`
	Channels int64 `json:"channels"`
}

type TopGamesResponse struct {
	Total    int32     `json:"_total"`
	TopGames []TopGame `json:"top"`
}

func (t Twitch) GetTopGames(page int32) (TopGamesResponse, error) {
	var gamesResponse TopGamesResponse

	resp, err := t.SendRequest("GET", "/games/top", "v3", t.BuildPageArgs(page))
	if err != nil {
		return gamesResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gamesResponse, err
	}

	err = json.Unmarshal(body, &gamesResponse)
	if err != nil {
		return gamesResponse, err
	}

	return gamesResponse, err
}
