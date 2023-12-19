package onosclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ParseIntent(body []byte) (Intent, error) {
	resp := Intent{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func ParseIntents(body []byte) (Intents, error) {
	resp := Intents{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *Client) GetIntents() (Intents, error) {
	resp := Intents{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/intents?detail=true", c.HostURL), nil)
	if err != nil {
		return resp, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return resp, err
	}

	resp, err = ParseIntents(body)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) GetIntent(intent Intent) (Intent, error) {
	resp := Intent{}
	if intent.AppID == "" || intent.Key == "" {
		return resp, errors.New("invalid intent; must include AppID, Key")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/intents/%s/%s", c.HostURL, intent.AppID, intent.Key), nil)
	req.Header.Add("Accept", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return resp, err
	}

	resp, err = ParseIntent(body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) CreateIntent(intent Intent) (Intent, error) {
	resp := Intent{}
	if intent.AppID == "" || intent.Type == "" || intent.One == "" || intent.Two == "" {
		return resp, errors.New("invalid intent; must include AppID, Type, One, Two")
	}

	rb, err := json.Marshal(intent)
	if err != nil {
		return resp, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/intents", c.HostURL), strings.NewReader(string(rb)))
	_, err = c.doRequest(req)
	if err != nil {
		return resp, err
	}

	resp, err = c.GetIntent(intent)
	attempts := 0
	for err != nil {
		if attempts >= 5 {
			break
		}
		time.Sleep(250 * time.Millisecond)
		resp, err = c.GetIntent(intent)
	}
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) UpdateIntent(intent Intent) (Intent, error) {
	//you can simply call create intent instead.
	//as long as the key already exists, the intent will be updated
	//this just runs a get to confirm the intent exists first in case an update only is needed.
	resp := Intent{}
	if intent.AppID == "" || intent.Type == "" || intent.One == "" || intent.Two == "" || intent.Key == "" {
		return resp, errors.New("invalid intent; must include AppID, Type, One, Two, Key")
	}

	_, err := c.GetIntent(intent)
	if err != nil {
		return resp, err
	}

	resp, err = c.CreateIntent(intent)
	if err != nil {
		return resp, err
	}

	//sometimes the original valuse are still returned if a get is ran very quickly after the update.
	//It seems the values take a little time to update. Added a wait and loop to retry
	attempts := 0
	for resp.One != intent.One || resp.Two != intent.Two || resp.Key != intent.Key || resp.AppID != intent.AppID {
		if attempts >= 5 {
			break
		}
		time.Sleep(250 * time.Millisecond)
		resp, err = c.GetIntent(intent)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func (c *Client) DeleteIntent(intent Intent) error {
	//this should return 200 for success and 204 for failure (no content), but onos api currently always returns 204 so there's no way to check the success/failure besides running another get and comparing.

	if intent.AppID == "" || intent.Key == "" {
		return errors.New("invalid intent; must include AppID, Key")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/intents/%s/%s", c.HostURL, intent.AppID, intent.Key), nil)
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
