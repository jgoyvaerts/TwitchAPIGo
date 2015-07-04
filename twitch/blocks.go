package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Block struct {
	User      User      `json=:"user"`
	UpdatedAt time.Time `json=:"updated_at"`
	ID        int64     `json=:"_id"`
}

func (t Twitch) GetUserBlockList(user string, page int32) ([]Block, error) {
	var blocks []Block
	req := "/users/" + user + "/blocks"

	resp, err := t.SendRequest("GET", req, "v3", t.BuildPageArgs(page))
	if err != nil {
		return blocks, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return blocks, err
	}

	err = json.Unmarshal(body, &blocks)
	if err != nil {
		return blocks, err
	}

	return blocks, err
}

func (t Twitch) AddUserToBlockList(user string, targetUser string) (Block, error) {
	var block Block
	req := "/users/" + user + "/blocks/" + targetUser

	resp, err := t.SendRequest("PUT", req, "v3", Options{})
	if err != nil {
		return block, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return block, err
	}

	err = json.Unmarshal(body, &block)
	if err != nil {
		return block, err
	}

	return block, err
}

func (t Twitch) RemoveUserFromBlockList(user string, targetUser string) (bool, error) {

	req := "/users/" + user + "/blocks/" + targetUser

	resp, err := t.SendRequest("DELETE", req, "v3", Options{})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return false, err
	}

	return true, err
}
