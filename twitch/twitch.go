package twitch

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	API_BASE_URL = "https://api.twitch.tv/kraken"
)

type Twitch struct {
	ClientID     string
	ItemsPerPage int32
	OAuthToken   string
	debug        bool
	out          io.Writer
	writer       io.Writer
}

type Options map[string]interface{}

func (o Options) String() string {
	if len(o) == 0 {
		return ""
	}
	out := "?"
	count := 1
	for k, v := range o {
		value := fmt.Sprintf("%v", v)
		out += fmt.Sprint(k, "=", url.QueryEscape(value))
		if count != len(o) {
			out += "&"
		}
		count += 1
	}
	return out
}

func NewClient(ClientID string) *Twitch {
	t := new(Twitch)
	t.ClientID = ClientID
	t.ItemsPerPage = 25
	t.debug = false
	t.out = ioutil.Discard
	t.writer = os.Stdout
	return t
}

func (t *Twitch) SetDebug(debug bool) {
	t.debug = debug
	t.updateWriter()
}

func (t *Twitch) SetDebugOutput(writer io.Writer) {
	t.writer = writer
	t.updateWriter()
}

func (t *Twitch) updateWriter() {
	if t.debug {
		t.out = t.writer
	} else {
		t.out = ioutil.Discard
	}
}

func (t Twitch) SendRequest(requestType, request, APIVersion string, options Options) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(requestType, API_BASE_URL+request+options.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.twitchtv."+APIVersion+"+json")

	if &t.OAuthToken != nil {
		req.Header.Add("Authorization", "OAuth "+t.OAuthToken)
	}

	if &t.ClientID != nil {
		req.Header.Add("Client-ID", t.ClientID)
	}

	fmt.Fprintf(t.out, "Request: %+v\n", req)

	response, err := client.Do(req)

	fmt.Fprintf(t.out, "Response: %+v\n", response)

	if err != nil {
		return response, err
	}

	if response.Status[0] == '4' || response.Status[0] == '5' {
		err = errors.New(response.Status)
	}

	return response, err

}

func (t Twitch) BuildPageArgs(page int32) Options {
	if page < 1 {
		page = 1
	}
	options := make(Options)
	options["limit"] = t.ItemsPerPage
	options["offset"] = (page - 1) * t.ItemsPerPage
	return options
}

func (t Twitch) MergeOptions(a Options, b Options) Options {
	for k, v := range a {
		b[k] = v
	}

	return b
}
