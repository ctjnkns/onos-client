package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetIntents() (Intents, error) {
	intents := Intents{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/intents?detail=true", c.HostURL), nil)
	body, err := c.doRequest(req)
	if err != nil {
		return intents, err
	}

	//fmt.Println("String:", string(body))

	err = json.Unmarshal(body, &intents)
	if err != nil {
		return intents, err
	}

	//fmt.Println("Go:", intents.Intent[0].Type)

	return intents, nil
}

func (c *Client) CreateIntent(intent Intent) error {
	rb, err := json.Marshal(intent)
	if err != nil {
		return err
	}

	fmt.Println(string(rb))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/intents", c.HostURL), strings.NewReader(string(rb)))
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteIntent(intent Intent) error {
	//this should return 200 for success and 204 for failure (no content), but onos api currently always returns 204 so there's no way to check the success/failure besides running another get and comparing.

	fmt.Println(intent.AppID)
	fmt.Println(intent.Key)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/intents/%s/%s", c.HostURL, intent.AppID, intent.Key), nil)
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
