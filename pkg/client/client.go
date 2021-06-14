package client

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	statusUrl            string
	statusUpdateInterval time.Duration
	httpClient           *http.Client

	latestStatus       *Status
	latestStatusUpdate time.Time
}

const (
	DefaultStatusURL            string        = "https://status.vatsim.net/status.json"
	DefaultStatusUpdateInterval time.Duration = 10 * time.Minute
)

type Option func(*Client)

func WithStatusURL(url string) Option {
	return func(c *Client) {
		c.statusUrl = url
	}
}

func WithStatusUpdateInterval(d time.Duration) Option {
	return func(c *Client) {
		c.statusUpdateInterval = d
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		statusUrl:            DefaultStatusURL,
		statusUpdateInterval: DefaultStatusUpdateInterval,
		httpClient:           new(http.Client),
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *Client) PullData() (Response, error) {
	s, err := c.status()
	if err != nil {
		return Response{}, err
	}

	resp, err := c.httpClient.Get(s.Data.V3[0])
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	r := Response{}
	err = json.NewDecoder(resp.Body).Decode(&r)

	return r, err
}

func (c *Client) status() (Status, error) {
	if c.latestStatus != nil && c.latestStatusUpdate.Add(-1*c.statusUpdateInterval).After(time.Now()) {
		return *c.latestStatus, nil
	}

	resp, err := c.httpClient.Get(c.statusUrl)
	if err != nil {
		return Status{}, err
	}
	defer resp.Body.Close()

	s := Status{}
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return s, err
	}

	c.latestStatus = &s
	c.latestStatusUpdate = time.Now()

	return s, nil
}
