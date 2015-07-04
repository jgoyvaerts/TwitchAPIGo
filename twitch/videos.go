package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Video struct {
	Title         string    `json:"title"`
	Descripion    string    `json:"description"`
	BroadcastID   int64     `json:"broadcast_id"`
	Status        string    `json:"status"`
	RecordedAt    time.Time `json:"recorded_at"`
	Game          string    `json:"game"`
	ID            string    `json:"_id"`
	TagList       string    `json:"tag_list"`
	Length        int32     `json:"length"`
	Preview       string    `json:"preview"`
	URL           string    `json:"url"`
	Views         int64     `json:"views"`
	BroadcastType string    `json:"broadcast_type"`
	Channel       Channel   `json:"channel"`
}

type ChannelVideosResponse struct {
	Total  int32   `json:"_total"`
	Videos []Video `json:"videos"`
}

func (t Twitch) GetVideo(videoID string) (Video, error) {
	var video Video

	resp, err := t.SendRequest("GET", "/videos/"+videoID, "v3", Options{})
	if err != nil {
		return video, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return video, err
	}

	err = json.Unmarshal(body, &video)
	return video, err
}

func (t Twitch) GetTopVideos(game string, period string, page int32) ([]Video, error) {
	var videos []Video

	options := make(Options)

	options["game"] = game
	if period == "" {
		period = "week"
	}
	options["period"] = "week"

	resp, err := t.SendRequest("GET", "/videos/top", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
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

func (t Twitch) GetChannelVideos(channel string, broadcasts bool, hls bool, page int32) (ChannelVideosResponse, error) {
	var channelVideosResponse ChannelVideosResponse

	options := make(Options)

	options["broadcasts"] = broadcasts
	options["hls"] = hls

	resp, err := t.SendRequest("GET", "/channels/"+channel+"/videos", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
	if err != nil {
		return channelVideosResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channelVideosResponse, err
	}

	err = json.Unmarshal(body, &channelVideosResponse)
	return channelVideosResponse, err
}
