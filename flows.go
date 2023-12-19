package onosclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetFlows() (Flows, error) {
	flows := Flows{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flows", c.HostURL), nil)
	if err != nil {
		return flows, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return flows, err
	}

	//fmt.Println("String:", string(body))

	err = json.Unmarshal(body, &flows)
	if err != nil {
		return flows, err
	}

	//fmt.Println("Go:", flows.Intent[0].Type)

	return flows, nil
}
