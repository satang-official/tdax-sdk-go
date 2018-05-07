package client

type Client struct {
	url       string
	userID    string
	apiKey    string
	apiSecret string
}

// const defaultURL = "https://api.tdax.com/"

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
