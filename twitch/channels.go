package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Channel struct {
	Game                string    `json:"game"`
	Name                string    `json:"name"`
	StreamKey           string    `json:"stream_key"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Teams               []Team    `json:"teams"`
	Banner              string    `json:"banner"`
	VideoBanner         string    `json:"video_banner"`
	Background          string    `json:"background"`
	Logo                string    `json:"logo"`
	ID                  int64     `json:"_id"`
	Mature              bool      `json:"mature"`
	Login               string    `json:"login"`
	URL                 string    `json:"url"`
	Email               string    `json:"email"`
	Status              string    `json:"status"`
	BroadcasterLanguage string    `json:"broadcaster_language"`
	DisplayName         string    `json:"display_name"`
	Delay               int32     `json:"delay"`
	Language            string    `json:"language"`
	ProfileBanner       string    `json:"profile_banner"`
	ProfileBannerColor  string    `json:"profile_banner_background_color"`
	Partner             bool      `json:"partner"`
	Views               int64     `json:"views"`
	Followers           int64     `json:"followers"`
}

func (t Twitch) GetChannel(channelName string) (Channel, error) {
	var channel Channel

	resp, err := t.SendRequest("GET", "/channels/"+channelName, "v3", Options{})
	if err != nil {
		return channel, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channel, err
	}

	err = json.Unmarshal(body, &channel)
	return channel, err
}

func (t Twitch) GetOwnChannel() (Channel, error) {
	var channel Channel

	resp, err := t.SendRequest("GET", "/channel", "v3", Options{})
	if err != nil {
		return channel, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channel, err
	}

	err = json.Unmarshal(body, &channel)
	return channel, err
}

func (t Twitch) GetChannelEditors(channelName string) ([]User, error) {
	var users []User

	resp, err := t.SendRequest("GET", "/channels/"+channelName+"/editors", "v3", Options{})
	if err != nil {
		return users, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return users, err
	}

	type allUsers struct {
		Users []User `json:"users"`
	}

	var all allUsers

	err = json.Unmarshal(body, &all)
	return all.Users, err
}

func (t Twitch) GetChannelTeams(channelName string) ([]Team, error) {
	var teams []Team

	resp, err := t.SendRequest("GET", "/channels/"+channelName+"/teams", "v3", Options{})
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
