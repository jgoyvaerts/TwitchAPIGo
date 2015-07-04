package twitch

import (
	"encoding/json"
	"io/ioutil"
)

type ChannelsResponse struct {
	Channels []Channel `json:"channels"`
	Total    int32     `json:"_total"`
}

func (t Twitch) SearchChannels(query string, page int32) (ChannelsResponse, error) {
	var channelsResponse ChannelsResponse
	options := make(Options)
	options["q"] = query

	resp, err := t.SendRequest("GET", "/search/channels", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
	if err != nil {
		return channelsResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channelsResponse, err
	}

	err = json.Unmarshal(body, &channelsResponse)
	return channelsResponse, err
}

func (t Twitch) SearchStreams(query string, hls bool, page int32) (StreamsResponse, error) {
	var streamsResponse StreamsResponse
	options := make(Options)
	options["q"] = query
	options["hls"] = hls

	resp, err := t.SendRequest("GET", "/search/streams", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
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

func (t Twitch) SearchGames(query string, live bool) ([]Game, error) {
	var games []Game
	options := make(Options)
	options["q"] = query
	options["type"] = "suggest"
	options["live"] = live

	resp, err := t.SendRequest("GET", "/search/games", "v3", options)
	if err != nil {
		return games, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return games, err
	}

	type allGames struct {
		Games []Game `json:"games"`
	}
	var all allGames

	err = json.Unmarshal(body, &all)
	return all.Games, err
}
