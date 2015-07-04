package twitch

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

type Stream struct {
	Game      string    `json:"game"`
	Channel   Channel   `json:"channel"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Viewers   int64     `json:"viewers"`
	Preview   Preview   `json:"preview"`
}

type StreamsResponse struct {
	Streams []Stream `json:"streams"`
	Total   int32    `json:"_total"`
}

type FeaturedStream struct {
	Stream    Stream `json:"stream"`
	Scheduled bool   `json:"scheduled"`
	Sponsored bool   `json:"sponsored"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Image     string `json:"image"`
}

type FeaturedStreamsResponse struct {
	FeaturedStreams []FeaturedStream `json:"featured"`
}

type StreamsSummaryResponse struct {
	Viewers  int64 `json:"viewers"`
	Channels int32 `json:"channels"`
}

type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

func (t Twitch) GetStreams(gameName string, channels []string, page int32) (StreamsResponse, error) {
	var streamsResponse StreamsResponse
	options := make(Options)
	if gameName != "" {
		options["game"] = gameName
	}
	if len(channels) > 0 {
		options["channel"] = strings.Join(channels, ",")
	}
	resp, err := t.SendRequest("GET", "/streams", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
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

func (t Twitch) GetFeaturedStreams(page int32) (FeaturedStreamsResponse, error) {
	var featuredStreamsResponse FeaturedStreamsResponse
	resp, err := t.SendRequest("GET", "/streams/featured", "v3", t.BuildPageArgs(page))
	if err != nil {
		return featuredStreamsResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return featuredStreamsResponse, err
	}

	err = json.Unmarshal(body, &featuredStreamsResponse)
	return featuredStreamsResponse, err
}

func (t Twitch) GetStreamsSummary(gameName string) (StreamsSummaryResponse, error) {
	var streamsSummaryResponse StreamsSummaryResponse
	options := make(Options)
	if gameName != "" {
		options["game"] = gameName
	}
	resp, err := t.SendRequest("GET", "/streams/summary", "v3", options)
	if err != nil {
		return streamsSummaryResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return streamsSummaryResponse, err
	}

	err = json.Unmarshal(body, &streamsSummaryResponse)
	return streamsSummaryResponse, err
}
