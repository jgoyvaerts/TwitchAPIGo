package twitch

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Subscription struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"_id"`
	User      User      `json:"user"`
}

type ChannelSubscription struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"_id"`
	Channel   Channel   `json:"channel"`
}

type ChannelSubscriptionsResponse struct {
	Total         int32          `json:"_total"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (t Twitch) GetChannelSubscriptions(channelName string, ascending bool, page int32) (ChannelSubscriptionsResponse, error) {
	var channelSubscriptionResponse ChannelSubscriptionsResponse

	options := make(Options)
	if ascending {
		options["direction"] = "asc"
	} else {
		options["direction"] = "desc"
	}

	resp, err := t.SendRequest("GET", "/channels/"+channelName+"/subscriptions", "v3", t.MergeOptions(options, t.BuildPageArgs(page)))
	if err != nil {
		return channelSubscriptionResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return channelSubscriptionResponse, err
	}

	err = json.Unmarshal(body, &channelSubscriptionResponse)
	return channelSubscriptionResponse, err
}

func (t Twitch) GetChannelHasSubscriber(channelName string, userName string) (Subscription, error) {
	var subscription Subscription

	resp, err := t.SendRequest("GET", "/channels/"+channelName+"/subscriptions/"+userName, "v3", Options{})
	if err != nil {
		return subscription, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return subscription, err
	}

	err = json.Unmarshal(body, &subscription)
	return subscription, err
}

func (t Twitch) GetUserIsSubscribed(userName string, channelName string) (ChannelSubscription, error) {
	var subscription ChannelSubscription

	resp, err := t.SendRequest("GET", "/users/"+userName+"/subscriptions/"+channelName, "v3", Options{})
	if err != nil {
		return subscription, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return subscription, err
	}

	err = json.Unmarshal(body, &subscription)
	return subscription, err
}
