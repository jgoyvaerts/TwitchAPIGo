package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type User struct {
	Type          string          `json:"type"`
	Logo          string          `json:"logo"`
	Name          string          `json:"name"`
	DisplayName   string          `json:"display_name"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	ID            int64           `json:"_id"`
	Bio           string          `json:"bio"`
	Partnered     bool            `json:"partnered"`
	Email         string          `json:"email"`
	Notifications map[string]bool `json:"notifications"`
}

func (t Twitch) GetUser(userName string) (User, error) {
	var user User

	resp, err := t.SendRequest("GET", "/users/"+userName, "v3", Options{})
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	return user, err
}

func (t Twitch) GetOwnUser() (User, error) {
	var user User

	resp, err := t.SendRequest("GET", "/user", "v3", Options{})
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	return user, err
}

func (t Twitch) GetStreamsFollowed(page int32) (StreamsResponse, error) {
	var streamsResponse StreamsResponse

	resp, err := t.SendRequest("GET", "/streams/followed", "v3", t.BuildPageArgs(page))
	if err != nil {
		return streamsResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return streamsResponse, err
	}

	err = json.Unmarshal(body, &streamsResponse)
	return streamsResponse, err
}

func (t Twitch) GetVideosFollowed(page int32) ([]Video, error) {
	var videos []Video

	resp, err := t.SendRequest("GET", "/videos/followed", "v3", t.BuildPageArgs(page))
	if err != nil {
		return videos, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return videos, err
	}

	type all struct {
		Videos []Video `json:"videos"`
	}
	var allVideos all

	err = json.Unmarshal(body, &allVideos)
	return allVideos.Videos, err
}
