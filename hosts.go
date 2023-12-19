package onosclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseHosts(body []byte) (Hosts, error) {
	resp := Hosts{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *Client) GetHosts() (Hosts, error) {
	resp := Hosts{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/hosts", c.HostURL), nil)
	if err != nil {
		return resp, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return resp, err
	}

	resp, err = ParseHosts(body)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
