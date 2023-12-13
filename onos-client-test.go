package main

import (
	"fmt"
	"io"
	"net/http"
	"onosclient"
)

const HostURL string = "http://localhost:8181/onos/v1/"

func main() {
	username := "onos"
	password := "rocks"

	client, err := onosclient.NewClient(HostURL, username, password)
	if err != nil {
		fmt.Errorf(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/intents", client.HostURL), nil)
	req.SetBasicAuth(username, password)
	res, err := client.doRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("status: %d, body: %s", res.StatusCode, body)
	}

	fmt.Println("Body Len:", len(body))
	fmt.Println("String intent:")
	s := string(body)
	fmt.Println(s)

}
