package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Token struct {
	Authorization Authorization `json:"authorization"`
	UserName      string        `json:"user_name"`
	Valid         bool          `json:"valid"`
}

type Authorization struct {
	Scopes    []string  `json:"scopes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Twitch) GetRoot() (Token, error) {
	var token Token

	resp, err := t.SendRequest("GET", "/", "v3", Options{})
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return token, err
	}

	type respToken struct {
		Token Token `json:"token"`
	}
	var all respToken

	err = json.Unmarshal(body, &all)
	return all.Token, err
}
