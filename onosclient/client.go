package onosclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Username   string
	Password   string
}

// NewClient
func NewClient(host, username string, password string) (*Client, error) {
	//Set defaults
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    host,
		Username:   username,
		Password:   password,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Username, c.Password)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
