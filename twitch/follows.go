package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Follower struct {
	User          User      `json:"user"`
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
}

type ChannelFollow struct {
	Channel       Channel   `json:"channel"`
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
}

type FollowersResponse struct {
	Total     int32      `json:"_total"`
	Followers []Follower `json:"follows"`
}

type FollowingResponse struct {
	Total          int32           `json:"_total"`
	ChannelFollows []ChannelFollow `json:"follows"`
}

func (t Twitch) GetChannelFollowers(channel string, ascending bool, page int32) (FollowersResponse, error) {
	var followersResponse FollowersResponse
	req := "/channels/" + channel + "/follows"

	options := make(Options)

	if ascending {
		options["direction"] = "asc"
	} else {
		options["direction"] = "desc"
	}

	resp, err := t.SendRequest("GET", req, "v3", t.MergeOptions(options, t.BuildPageArgs(page)))

	if err != nil {
		return followersResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return followersResponse, err
	}

	err = json.Unmarshal(body, &followersResponse)
	if err != nil {
		return followersResponse, err
	}

	return followersResponse, err
}

func (t Twitch) GetChannelsFollowing(user string, ascending bool, sortBy string, page int32) (FollowingResponse, error) {
	var followingResponse FollowingResponse
	req := "/users/" + user + "/follows/channels"

	options := make(Options)

	if ascending {
		options["direction"] = "asc"
	} else {
		options["direction"] = "desc"
	}

	if sortBy == "" {
		sortBy = "created_at"
	}
	options["sortby"] = sortBy

	resp, err := t.SendRequest("GET", req, "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
	if err != nil {
		return followingResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return followingResponse, err
	}

	err = json.Unmarshal(body, &followingResponse)
	if err != nil {
		return followingResponse, err
	}

	return followingResponse, err
}

func (t Twitch) GetIsUserFollowingChannel(user string, channel string) (bool, ChannelFollow, error) {
	var channelFollow ChannelFollow
	req := "/users/" + user + "/follows/channels/" + channel

	resp, err := t.SendRequest("GET", req, "v3", Options{})
	if err != nil {
		return false, channelFollow, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return false, channelFollow, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, channelFollow, err
	}

	err = json.Unmarshal(body, &channelFollow)
	if err != nil {
		return false, channelFollow, err
	}

	return true, channelFollow, err
}

func (t Twitch) AddUserToFollowers(user string, channel string, notifications bool) (ChannelFollow, error) {
	var channelFollow ChannelFollow
	req := "/users/" + user + "/follows/channels/" + channel

	options := make(Options)
	options["notifications"] = notifications

	resp, err := t.SendRequest("PUT", req, "v3", options)
	if err != nil {
		return channelFollow, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channelFollow, err
	}

	err = json.Unmarshal(body, &channelFollow)
	if err != nil {
		return channelFollow, err
	}

	return channelFollow, err
}

func (t Twitch) RemoveUserFromFollowers(user string, channel string) (bool, error) {

	req := "/users/" + user + "/follows/channels/" + channel

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
