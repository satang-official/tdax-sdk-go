package client

import (
	"fmt"
	"net/http"

	resty "gopkg.in/resty.v1"
)

type Client struct {
	url       string
	userID    string
	apiKey    string
	apiSecret string
}

type Error struct {
	StatusCode int
	Response   []byte
}

func (e Error) Error() string {
	return fmt.Sprintf("%d:%s", e.StatusCode, string(e.Response))
}

func init() {
	resty.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if resp.StatusCode() >= http.StatusBadRequest {
			return Error{resp.StatusCode(), resp.Body()}
		}
		return nil // if its success otherwise return error
	})
}

func NewClient(url, userID, apiKey, apiSecret string) Client {
	return Client{
		url:       url,
		userID:    userID,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (c Client) URL() string {
	return c.url
}

func (c Client) UserID() string {
	return c.userID
}

func (c Client) APIKey() string {
	return "TDAX-API " + c.apiKey
}

func (c Client) APISecret() string {
	return c.apiSecret
}
